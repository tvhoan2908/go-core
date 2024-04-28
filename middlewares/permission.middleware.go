package middlewares

import (
	"go-core/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func PermissionMiddleware(permission string) gin.HandlerFunc {
	return func(c *gin.Context) {
		userId, _ := c.Get("UserID")
		issuedAt, _ := c.Get("IssuedAt")

		user, err := services.ValidateUserFromToken(userId.(uint64), issuedAt.(int64))
		if err != nil {
			c.Error(NewError(http.StatusUnauthorized, "Unauthorized."))
			c.Abort()
			return
		}

		if services.InitUser(user).HasPermission(permission) {
			c.Next()
			return
		}

		c.Error(NewError(403, "You can not access this page."))
		c.Abort()
	}
}
