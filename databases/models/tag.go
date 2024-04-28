package models

import "time"

type Tag struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"not null;type:varchar(255)"`
	Slug      string `gorm:"not null;type:varchar(255);unique"`
	CreatedAt time.Time
}
