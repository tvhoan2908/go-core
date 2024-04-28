package services

import (
	"go-core/databases/models"
	"go-core/types"

	"gorm.io/gorm"
)

type CategoryService interface {
	CreateCategory(request *types.StoreCategoryDTO, userId uint64) (uint, error)
	UpdateCategory(request *types.StoreCategoryDTO, id uint64) error
	CategoryList() ([]*types.CategoryDTO, error)
}

type categoryService struct {
	db *gorm.DB
}

func NewCategoryService(db *gorm.DB) CategoryService {
	return &categoryService{db: db}
}

func (s *categoryService) CreateCategory(request *types.StoreCategoryDTO, userId uint64) (uint, error) {
	category := models.Category{
		Name:        request.Name,
		Slug:        request.Slug,
		Description: request.Description,
		UserID:      &userId,
		ParentID:    request.ParentID,
	}

	result := s.db.Create(&category)

	return category.ID, result.Error
}

func (s *categoryService) UpdateCategory(request *types.StoreCategoryDTO, id uint64) error {
	var category models.Category
	categoryRs := s.db.First(&category, id)
	if categoryRs.Error != nil {
		return categoryRs.Error
	}

	category.Name = request.Name
	category.Description = request.Description
	category.Slug = request.Slug
	category.ParentID = request.ParentID

	rs := s.db.Save(&category)

	return rs.Error
}

func (s *categoryService) CategoryList() ([]*types.CategoryDTO, error) {
	var responses []*types.CategoryDTO

	result := s.db.Model(&models.Category{}).Order("id DESC").Preload("Parent", func(db *gorm.DB) *gorm.DB {
		return db.Model(&models.Category{}).Find(&types.BaseRelationDTO{})
	}).Limit(20).Offset(0).Find(&responses)
	if result.Error != nil {
		return nil, result.Error
	}

	return responses, nil
}
