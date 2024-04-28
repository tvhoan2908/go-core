package handlers

import (
	Middlewares "go-core/middlewares"
	"go-core/services"
	"go-core/types"
	Utils "go-core/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserHandler interface {
	StoreUser(c *gin.Context)
	UpdateUser(c *gin.Context)
	ChangePassword(c *gin.Context)
	UserList(c *gin.Context)
}

type userHandler struct {
	userService services.UserService
}

func NewUserHandler(service services.UserService) UserHandler {
	return &userHandler{userService: service}
}

// User godoc
// @Summary      	Get user logged info
// @Description  	Get user logged info
// @Tags         	User
// @Router       	/api/v1/admin/users [get]
// @Success				200 {object} handlers.BaseOutput{data=[]types.UserInfoDTO} "desc"
// @Security			Bearer
func UserInfo(c *gin.Context) {
	userId, _ := c.Get("UserID")

	response, err := services.UserInfo(userId.(uint64))
	if err != nil {
		c.Error(err)
		return
	}

	InitContext(c).ResponseEntity(&BaseOutput{
		Data: response,
	})
}

// User godoc
// @Summary      	Create user
// @Description  	Create user
// @Tags         	User
// @Param				 	Body body types.StoreUserRequest true "User"
// @Router       	/api/v1/admin/users [post]
// @Success				200 {object} handlers.BaseOutput{} "desc"
// @Security			Bearer
func (h *userHandler) StoreUser(c *gin.Context) {
	var request types.StoreUserRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.Error(&Middlewares.CustomError{
			Code:   http.StatusBadRequest,
			Errors: Utils.HandleErrorsMessage(err),
		})
		return
	}

	err := h.userService.StoreUser(&request)
	if err != nil {
		c.Error(err)
		return
	}

	InitContext(c).ResponseEntity(&BaseOutput{
		Code: http.StatusCreated,
	})
}

// User godoc
// @Summary      	Update user
// @Description  	Update user
// @Tags         	User
// @Param				 	Body body types.UpdateUserRequest true "User"
// @Router       	/api/v1/admin/users/{id} [put]
// @Success				200 {object} handlers.BaseOutput{} "desc"
// @Security			Bearer
func (h *userHandler) UpdateUser(c *gin.Context) {
	var request types.UpdateUserRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.Error(&Middlewares.CustomError{
			Code:   http.StatusBadRequest,
			Errors: Utils.HandleErrorsMessage(err),
		})
		return
	}

	userId, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.Error(err)
		return
	}

	if err := h.userService.UpdateUser(&request, userId); err != nil {
		c.Error(err)
		return
	}

	InitContext(c).ResponseEntity(&BaseOutput{})
}

// User godoc
// @Summary      	Change user password
// @Description  	Change user password
// @Tags         	User
// @Param				 	Body body types.ChangePasswordRequest true "User"
// @Router       	/api/v1/admin/users/change-password [patch]
// @Success				200 {object} handlers.BaseOutput{} "desc"
// @Security			Bearer
func (h *userHandler) ChangePassword(c *gin.Context) {
	var request types.ChangePasswordRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.Error(&Middlewares.CustomError{
			Code:   http.StatusBadRequest,
			Errors: Utils.HandleErrorsMessage(err),
		})
		return
	}

	userId, _ := c.Get("UserID")
	err := h.userService.ChangePassword(userId.(uint64), &request)
	if err != nil {
		c.Error(err)
		return
	}

	InitContext(c).ResponseEntity(&BaseOutput{})
}

// User godoc
// @Summary      	Get user list
// @Description  	Get user list
// @Tags         	User
// @Router       	/api/v1/admin/users [get]
// @Success				200 {object} handlers.BaseOutput{data=[]types.UserDTO} "desc"
// @Security			Bearer
func (h *userHandler) UserList(c *gin.Context) {
	responses, err := h.userService.UserList()
	if err != nil {
		c.Error(err)
		return
	}

	InitContext(c).ResponseEntity(&BaseOutput{
		Data: responses,
	})
}
