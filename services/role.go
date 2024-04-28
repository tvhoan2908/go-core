package services

import (
	"fmt"
	"go-core/config"
	"go-core/databases/models"
	"go-core/types"
	"go-core/utils"

	"gorm.io/gorm"
)

type RoleService interface {
	RoleList() ([]*types.RoleDTO, error)
	CreateRole(request *types.CreateRoleDTO, userId uint) (uint, error)
	UpdateRole(request *types.CreateRoleDTO, id uint) error
}

type roleService struct {
	db *gorm.DB
}

func NewRoleService(db *gorm.DB) RoleService {
	return &roleService{db: db}
}

func (s *roleService) RoleList() ([]*types.RoleDTO, error) {
	var responses []*types.RoleDTO
	result := s.db.Model(models.Role{}).Order("id DESC").Find(&responses)
	if result.Error != nil {
		return nil, result.Error
	}

	return responses, nil
}

func (s *roleService) CreateRole(request *types.CreateRoleDTO, userId uint) (uint, error) {
	db := s.db
	var user models.User
	userResult := db.First(&user, userId)
	if userResult.Error != nil {
		return 0, userResult.Error
	}

	var permissions []*models.Permission
	permissionResult := db.Find(&permissions, request.Permissions)
	if permissionResult.Error != nil {
		return 0, permissionResult.Error
	}

	role := models.Role{
		Name:        request.Name,
		Description: request.Description,
		User:        &user,
		Permissions: permissions,
	}

	result := db.Create(&role)

	return role.ID, result.Error
}

func (s *roleService) UpdateRole(request *types.CreateRoleDTO, id uint) error {
	db := s.db
	var role models.Role
	roleResult := db.First(&role, id)
	if roleResult.Error != nil {
		return roleResult.Error
	}

	var permissions []*models.Permission
	permissionResult := db.Find(&permissions, request.Permissions)
	if permissionResult.Error != nil {
		return permissionResult.Error
	}

	db.Model(&role).Association("Permissions").Clear()

	role.Name = request.Name
	role.Description = request.Description
	role.Permissions = permissions
	result := db.Save(&role)

	var usersID []*uint
	rs := db.Raw("SELECT user_id FROM user_roles WHERE role_id = ?", role.ID).Find(&usersID)
	if rs.Error == nil {
		for _, userId := range usersID {
			utils.DeleteRedisKey(fmt.Sprintf("%s%d", config.USER_CACHE_KEY, *userId))
		}
	}

	return result.Error
}
