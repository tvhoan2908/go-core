package models

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	Name        string `gorm:"not null;type:varchar(255)"`
	Slug        string `gorm:"not null;unique;type:varchar(255)"`
	Description *string
	Content     string
	UserID      uint
	User        User
	CategoryID  *int
	Category    *Category `gorm:"foreignKey:CategoryID"`
	Status      int       `gorm:"size:1;default:2;comment:1-Publish,2-Pending,3-UnPublish;index"`
	Tags        []*Tag    `gorm:"many2many:post_tags;"`
	MediaID     *uint
	Media       *Media
}
