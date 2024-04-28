package services

import (
	"go-core/databases/models"
	"go-core/types"
	"go-core/utils"

	"gorm.io/gorm"
)

type MediaService interface {
	StoreMedia(request *types.StoreMediaRequest) (*uint, error)
	MediaList(request *types.FilterMediaRequest) ([]*types.MediaDTO, error)
}

type mediaService struct {
	db *gorm.DB
}

func NewMediaService(db *gorm.DB) MediaService {
	return &mediaService{db: db}
}

func (s *mediaService) StoreMedia(request *types.StoreMediaRequest) (*uint, error) {
	db := s.db
	media := models.Media{
		FileName: request.FileName,
		FileMime: request.FileMime,
		FileSize: request.FileSize,
		FileType: request.FileType,
		UserID:   &request.UserID,
		Path:     request.Path,
	}

	rs := db.Create(&media)
	if rs.Error != nil {
		return nil, rs.Error
	}

	return &media.ID, nil
}

func (s *mediaService) MediaList(request *types.FilterMediaRequest) ([]*types.MediaDTO, error) {
	var resources []*types.MediaDTO

	db := s.db
	mediaFilter := models.Media{
		UserID: &request.UserID,
	}

	rs := db.Model(&models.Media{}).Where(&mediaFilter).Scopes(utils.Paginate(&request.BaseFilterRequest)).Order("id DESC").Find(&resources)
	if rs.Error != nil {
		return nil, rs.Error
	}

	return resources, nil
}
