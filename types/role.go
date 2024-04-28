package types

import "time"

type CreateRoleDTO struct {
	Name        string  `form:"name" json:"name" binding:"required"`
	Description *string `form:"description" json:"description"`
	Permissions []*int  `form:"permissions" json:"permissions" binding:"required"`
}

type RoleDTO struct {
	ID          uint      `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"createdAt"`
}
