package middlewares

import (
	Utils "go-core/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CustomError struct {
	Code   int         `json:"code"`
	Msg    string      `json:"message"`
	Errors interface{} `json:"errors,omitempty"`
}

func (e *CustomError) Error() string {
	return e.Msg
}

func NewError(Code int, Msg string) *CustomError {
	return &CustomError{
		Code: Code,
		Msg:  Msg,
	}
}

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next() // Call c.Next() to execute the middleware
		for _, e := range c.Errors {
			err := e.Err
			// Return code and message if there is a custom error
			if v, ok := err.(*CustomError); ok {
				if v.Msg == "" {
					v.Msg = "false"
				}
				c.JSON(v.Code, v)
			} else {
				c.JSON(http.StatusInternalServerError, &CustomError{
					Code: http.StatusInternalServerError,
					Msg:  err.Error(),
				})
			}

			Utils.LogError(c.Request, err)
		}
	}
}
