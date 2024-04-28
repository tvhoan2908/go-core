package types

import "mime/multipart"

type FilterMediaRequest struct {
	BaseFilterRequest
	UserID uint64
}

type StoreMediaRequest struct {
	FileName string
	FileMime string
	FileSize int64
	FileType int
	Path     string
	UserID   uint64
}

type MediaDTO struct {
	ID       uint   `json:"id"`
	FileName string `json:"fileName"`
	Path     string `json:"url"`
}

type UploadFileRequest struct {
	Files []*multipart.FileHeader `form:"files" binding:"required"`
}
