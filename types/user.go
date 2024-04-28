package types

import "time"

type ChangePasswordRequest struct {
	OldPassword string `json:"oldPassword" binding:"required"`
	NewPassword string `json:"newPassword" binding:"required"`
}

type StoreUserRequest struct {
	Username string  `json:"username" binding:"required,unique_name=users username"`
	Password string  `json:"password" binding:"required"`
	FullName string  `json:"fullName" binding:"required"`
	Email    *string `json:"email" binding:"required,unique_name=users email"`
	Roles    []uint  `json:"roles" binding:"required"`
}

type UpdateUserRequest struct {
	Password string `json:"password"`
	FullName string `json:"fullName" binding:"required"`
	Roles    []uint `json:"roles" binding:"required"`
}

type BaseUserDTO struct {
	ID          uint    `json:"id"`
	Username    string  `json:"username"`
	Email       *string `json:"email"`
	FullName    string  `json:"fullName"`
	Status      int     `json:"status"`
	AccountType int     `json:"accountType"`
}

type UserInfoDTO struct {
	BaseUserDTO
	Permissisons []*string `json:"permissions"`
}

type UserDTO struct {
	BaseUserDTO
	Roles     []*string `json:"roles"`
	CreatedAt time.Time `json:"created_at"`
}
