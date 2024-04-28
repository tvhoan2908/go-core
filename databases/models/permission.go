package models

import "gorm.io/gorm"

type Permission struct {
	gorm.Model
	Name        string `gorm:"not null;type:varchar(255);unique"`
	Description string
	ModuleID    *int
	Module      *Module
}
