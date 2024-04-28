package handlers

import (
	Middlewares "go-core/middlewares"
	"go-core/services"
	"go-core/types"
	Utils "go-core/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthHandler interface {
	Register(c *gin.Context)
	Login(c *gin.Context)
}

type authHandler struct {
	authService services.AuthService
}

func NewAuthHandler(service services.AuthService) AuthHandler {
	return &authHandler{authService: service}
}

// Auth godoc
// @Summary      	Register
// @Description  	Register system
// @Tags         	Auth
// @Param				 	Body body types.RegisterRequest true "Register"
// @Router       	/api/v1/auth/register [post]
// @Success       200  {object} handlers.BaseOutput{} "token"
func (h *authHandler) Register(c *gin.Context) {
	var request types.RegisterRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.Error(Middlewares.NewError(http.StatusBadRequest, err.Error()))
		return
	}

	password, _ := services.HashPassword(request.Password)
	request.Password = password

	user, err := h.authService.Register(request)
	if err != nil {
		c.Error(err)
		return
	}

	Utils.EmailChannel <- user
	InitContext(c).ResponseEntity(&BaseOutput{
		Message: "success",
	})
}

// Auth godoc
// @Summary      	Login
// @Description  	Login system
// @Tags         	Auth
// @Param				 	Body body types.LoginRequest true "Login"
// @Success       200  {object} handlers.BaseOutput{data=string} "token"
// @Router       	/api/v1/auth/login [post]
func (h *authHandler) Login(c *gin.Context) {
	var request types.LoginRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.Error(Middlewares.NewError(http.StatusBadRequest, err.Error()))
		return
	}

	user, err := h.authService.Login(request)
	if err != nil {
		c.Error(Middlewares.NewError(http.StatusUnauthorized, err.Error()))
		return
	}

	token, err := services.CreateToken(*user)
	if err != nil {
		c.Error(Middlewares.NewError(http.StatusUnprocessableEntity, err.Error()))
		return
	}

	InitContext(c).ResponseEntity(&BaseOutput{
		Data: token,
	})
}
