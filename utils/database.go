package utils

import (
	"go-core/types"

	"gorm.io/gorm"
)

func Paginate(r *types.BaseFilterRequest) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		page := r.Page
		if page <= 0 {
			page = 1
		}
		size := r.Size
		switch {
		case size > 100:
			size = 100
		case size <= 0:
			size = 20
		}

		offset := (page - 1) * size
		return db.Offset(offset).Limit(size)
	}
}
