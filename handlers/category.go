package handlers

import (
	Middlewares "go-core/middlewares"
	"go-core/services"
	"go-core/types"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CategoryHandler interface {
	CreateCategory(c *gin.Context)
	UpdateCategory(c *gin.Context)
	CategoryList(c *gin.Context)
}

type categoryHandler struct {
	categoryService services.CategoryService
}

func NewCategoryHandler(service services.CategoryService) CategoryHandler {
	return &categoryHandler{categoryService: service}
}

// Category godoc
// @Summary      	Create category
// @Description  	Create category
// @Tags         	Category
// @Param				 	Body body types.StoreCategoryDTO true "Register"
// @Router       	/api/v1/admin/categories [post]
// @Security			Bearer
// @Success       200  {object} handlers.BaseOutput{} "Category"
func (h *categoryHandler) CreateCategory(c *gin.Context) {
	var request types.StoreCategoryDTO
	if err := c.ShouldBindJSON(&request); err != nil {
		c.Error(Middlewares.NewError(http.StatusBadRequest, err.Error()))
		return
	}

	userId, _ := c.Get("UserID")
	insertedId, err := h.categoryService.CreateCategory(&request, userId.(uint64))
	if err != nil {
		c.Error(err)
		return
	}

	InitContext(c).ResponseEntity(&BaseOutput{
		Code: http.StatusCreated,
		Data: insertedId,
	})
}

// Category godoc
// @Summary      	Update category
// @Description  	Update category
// @Tags         	Category
// @Param				 	Body body types.StoreCategoryDTO true "Register"
// @Param					id path string true "CategoryID"
// @Router       	/api/v1/admin/categories/{id} [put]
// @Security			Bearer
// @Success       200  {object} handlers.BaseOutput{} "Category"
func (h *categoryHandler) UpdateCategory(c *gin.Context) {
	categoryId := c.Param("id")
	id, err := strconv.ParseUint(categoryId, 10, 64)
	if err != nil {
		c.Error(err)
		return
	}

	var request types.StoreCategoryDTO
	if err := c.ShouldBindJSON(&request); err != nil {
		c.Error(Middlewares.NewError(http.StatusBadRequest, err.Error()))
		return
	}

	if err := h.categoryService.UpdateCategory(&request, id); err != nil {
		c.Error(err)
		return
	}

	InitContext(c).ResponseEntity(&BaseOutput{})
}

// Category godoc
// @Summary      	Get category list
// @Description  	Get category list
// @Tags         	Category
// @Success				200 {object} BaseOutput{data=[]types.CategoryDTO} "desc"
// @Router       	/api/v1/admin/categories [get]
// @Security			Bearer
// @Success       200  {object} handlers.BaseOutput{data=[]types.CategoryDTO} "token"
func (h *categoryHandler) CategoryList(c *gin.Context) {
	responses, err := h.categoryService.CategoryList()
	if err != nil {
		c.Error(err)
		return
	}

	InitContext(c).ResponseEntity(&BaseOutput{
		Data: responses,
	})
}
