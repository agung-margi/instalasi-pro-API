package auth

import (
	"fmt"
	"net/http"
	"strings"

	"instalasi-pro/configs"
	helper "instalasi-pro/helpers"
	"instalasi-pro/modules/user"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type Service interface {
	AuthMiddleware() gin.HandlerFunc
}

type Claims struct {
	UserID    int    `json:"user_id"`
	RoleUser  string `json:"role"`
	ExpiresAt int64  `json:"exp"`
	jwt.StandardClaims
}

func AuthMiddleware() gin.HandlerFunc {
	var jwtKey = []byte(configs.AppConfig.JWTSecretKey)

	return func(c *gin.Context) {
		fmt.Println("AuthMiddleware called")
		token := c.Request.Header.Get("Authorization")
		if token == "" {
			response := helper.APIResponse("The request is unauthenticated.", http.StatusUnauthorized, "error", nil)
			c.JSON(http.StatusUnauthorized, response)
			c.Abort()
			return
		}

		parts := strings.Split(token, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			response := helper.APIResponse("Invalid token format.", http.StatusUnauthorized, "error", nil)
			c.JSON(http.StatusUnauthorized, response)
			c.Abort()
			return
		}

		token = parts[1]
		claims := &Claims{}

		_, err := jwt.ParseWithClaims(token, claims, func(t *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})
		if err != nil {
			response := helper.APIResponse("Invalid or expired token.", http.StatusUnauthorized, "error", nil)
			c.JSON(http.StatusUnauthorized, response)
			c.Abort()
			return
		}

		c.Set("currentUser", user.User{ID: claims.UserID, Role: claims.RoleUser})
		c.Next()
	}
}
