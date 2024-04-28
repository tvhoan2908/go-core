package handlers

import (
	"go-core/types"
	"net/http"

	"github.com/gin-gonic/gin"
)

type BaseOutput struct {
	Code       int                  `json:"code"`
	Message    string               `json:"message"`
	Data       interface{}          `json:"data"`
	Pagination *types.PaginationDTO `json:"pagination,omitempty"`
}

type NewContext struct {
	*gin.Context
}

func InitContext(c *gin.Context) *NewContext {
	return &NewContext{c}
}

func (c *NewContext) ResponseEntity(params *BaseOutput) {
	response := &BaseOutput{
		Code: http.StatusOK,
		Data: params.Data,
	}
	if params.Code > 0 {
		response.Code = params.Code
	}

	response.Message = "success"
	if params.Message != "" {
		response.Message = params.Message
	}

	if params.Pagination != nil {
		response.Pagination = params.Pagination
	}

	c.JSON(http.StatusOK, response)
}
