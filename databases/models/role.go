package models

import "gorm.io/gorm"

type Role struct {
	gorm.Model
	Name        string `gorm:"not null;type:varchar(255)"`
	Description *string
	UserID      *int
	User        *User
	Permissions []*Permission `gorm:"many2many:role_permissions;"`
}
