package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username       string  `gorm:"not null;unique;type:varchar(255);column:username"`
	Password       string  `gorm:"not null;column:password"`
	Email          *string `gorm:"unique;type:varchar(255)"`
	FullName       string
	Status         int       `gorm:"size:1;default:1;comment:1-Visible,2-Banned,3-Disabled"`
	AccountType    int       `gorm:"size:1;default:2;comment:1-Administrator,2-Normal Account"`
	Roles          []*Role   `gorm:"many2many:user_roles;"`
	TokenExpiredAt time.Time `gorm:"column:token_expired_at;index:user_idx_token_expired;comment:Thoi gian Token het han"`
}

// Override Table Name
func (User) TableName() string {
	return "users"
}
