package types

import "time"

type StoreCategoryDTO struct {
	Name        string  `form:"name" json:"name" binding:"required"`
	Slug        string  `form:"slug" json:"slug" binding:"required,unique_name=categories slug"`
	Description *string `form:"description" json:"description"`
	ParentID    *uint   `form:"parent_id" json:"parentId"`
}

type CategoryDTO struct {
	ID          uint             `json:"id"`
	Name        string           `json:"name"`
	Description string           `json:"description"`
	Slug        string           `json:"slug"`
	CreatedAt   time.Time        `json:"createdAt"`
	Parent      *BaseRelationDTO `json:"parent"`
	ParentID    *uint            `json:"parentId"`
}
