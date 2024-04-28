package models

import (
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	Name        string `gorm:"not null;type:varchar(255)"`
	Slug        string `gorm:"not null;type:varchar(255);unique"`
	Description *string
	UserID      *uint64
	User        *User
	ParentID    *uint
	Parent      *Category
}
