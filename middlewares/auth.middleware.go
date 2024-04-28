package middlewares

import (
	"fmt"
	"go-core/config"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

type Client struct {
	UserID   uint64
	Token    string
	IssuedAt int64
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		err, payload := TokenValid(c.Request)
		if err != nil {
			c.Error(NewError(http.StatusUnauthorized, err.Error()))
			c.Abort()
			return
		}

		c.Set("UserID", payload.UserID)
		c.Set("IssuedAt", payload.IssuedAt)
		c.Next()
	}
}

func TokenValid(r *http.Request) (error, *Client) {
	token, err := VerifyToken(r)
	if err != nil {
		return err, nil
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userId, err := strconv.ParseUint(fmt.Sprintf("%.f", claims["userId"]), 10, 64)
		issuedAt, err := strconv.ParseInt(fmt.Sprintf("%.f", claims["iat"]), 10, 64)
		if err != nil {
			return err, nil
		}

		return nil, &Client{UserID: userId, IssuedAt: issuedAt}
	}

	return err, nil
}

func VerifyToken(r *http.Request) (*jwt.Token, error) {
	tokenString := ExtractToken(r)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expected
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected singing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(config.SECRET_KEY), nil
	})

	if err != nil {
		return nil, err
	}

	return token, nil
}

func ExtractToken(r *http.Request) string {
	bearToken := r.Header.Get("Authorization")
	strArr := strings.Split(bearToken, " ")
	if len(strArr) == 2 {
		return strArr[1]
	}

	return ""
}
