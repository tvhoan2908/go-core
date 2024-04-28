package models

import "time"

type Media struct {
	ID        uint   `gorm:"primaryKey"`
	FileName  string `gorm:"not null;type:varchar(255)"`
	FileMime  string `gorm:"type:varchar(255)"`
	FileSize  int64
	FileType  int    `gorm:"size:1;default:1;comment:Loai File-Image,File,Video,Audio..."`
	Path      string `gorm:"not null"`
	UserID    *uint64
	User      *User
	CreatedAt time.Time
}
