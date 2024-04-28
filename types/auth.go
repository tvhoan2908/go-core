package types

type LoginRequest struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

type RegisterRequest struct {
	Username string  `form:"username" json:"username" binding:"required,unique_name=users username"`
	Password string  `form:"password" json:"password" binding:"required,min=6"`
	FullName string  `form:"full_name" json:"full_name" binding:"required"`
	Email    *string `form:"email" json:"email" binding:"unique_name=users email"`
}
