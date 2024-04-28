package handlers

import (
	Middlewares "go-core/middlewares"
	"go-core/services"
	"go-core/types"
	"go-core/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RoleHandler interface {
	RoleList(c *gin.Context)
	CreateRole(c *gin.Context)
	UpdateRole(c *gin.Context)
}

type roleHandler struct {
	roleService services.RoleService
}

func NewRoleHandler(service services.RoleService) RoleHandler {
	return &roleHandler{roleService: service}
}

// Role godoc
// @Summary      	Get role list
// @Description  	Get role list
// @Tags         	Role
// @Router       	/api/v1/admin/roles [get]
// @Security			Bearer
// @Success       200  {object} handlers.BaseOutput{data=[]types.RoleDTO} "Role"
func (s *roleHandler) RoleList(c *gin.Context) {
	responses, err := s.roleService.RoleList()
	if err != nil {
		c.Error(err)
		return
	}

	InitContext(c).ResponseEntity(&BaseOutput{
		Data: responses,
	})
}

// Role godoc
// @Summary      	Create role
// @Description  	Create role
// @Tags         	Role
// @Param				 	Body body types.CreateRoleDTO true "Role"
// @Router       	/api/v1/admin/roles [post]
// @Security			Bearer
// @Success       200  {object} handlers.BaseOutput{} "Role"
func (s *roleHandler) CreateRole(c *gin.Context) {
	var request types.CreateRoleDTO
	if err := c.ShouldBindJSON(&request); err != nil {
		c.Error(Middlewares.NewError(http.StatusBadRequest, err.Error()))
		return
	}

	userId, _ := c.Get("UserID")
	roleId, err := s.roleService.CreateRole(&request, userId.(uint))
	if err != nil {
		c.Error(err)
		return
	}

	InitContext(c).ResponseEntity(&BaseOutput{
		Code: http.StatusCreated,
		Data: roleId,
	})
}

// Role godoc
// @Summary      	Update role
// @Description  	Update role
// @Tags         	Role
// @Param				 	Body body types.CreateRoleDTO true "Role"
// @Param					id path string true "RoleID"
// @Router       	/api/v1/admin/roles/{id} [put]
// @Security			Bearer
// @Success       200  {object} handlers.BaseOutput{} "Role"
func (s *roleHandler) UpdateRole(c *gin.Context) {
	var request types.CreateRoleDTO
	if err := c.ShouldBindJSON(&request); err != nil {
		c.Error(Middlewares.NewError(http.StatusBadRequest, err.Error()))
		return
	}

	id := utils.ParseUInt(c.Param("id"))
	if err := s.roleService.UpdateRole(&request, id); err != nil {
		c.Error(err)
		return
	}

	InitContext(c).ResponseEntity(&BaseOutput{})
}
