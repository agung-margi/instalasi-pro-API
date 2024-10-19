package auth

import (
	"fmt"
	"net/http"
	"os"

	"strings"

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
	ExpiresAt int64  `json:"exp"` // Gunakan int64 untuk expiration
	jwt.StandardClaims
}

func AuthMiddleware() gin.HandlerFunc {
	var jwtKey = []byte(os.Getenv("JWT_SECRET_KEY"))

	fmt.Println("JWT_SECRET_KEY:", jwtKey)
	return func(c *gin.Context) {
		fmt.Println("AuthMiddleware called")
		token := c.Request.Header.Get("Authorization")
		fmt.Println("Token:", token)
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

		fmt.Println("Token:", token)
		claims := jwt.MapClaims{}
		_, err := jwt.ParseWithClaims(token, claims, func(t *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})
		if err != nil {
			fmt.Println("Error parsing token:", err) // Log error spesifik
			response := helper.APIResponse("Invalid or expired token.", http.StatusUnauthorized, "error", nil)
			c.JSON(http.StatusUnauthorized, response)
			c.Abort()
			return
		}

		// fmt.Println("Token Claims:", claims)
		fmt.Println("Token Claims:", claims["exp"])
		fmt.Println("Token Claims:", claims["role"])

		// if err != nil || !tkn.Valid {
		// 	response := helper.APIResponse("Invalid or expired token.", http.StatusUnauthorized, "error", nil)
		// 	c.JSON(http.StatusUnauthorized, response)
		// 	c.Abort()
		// 	return
		// }

		userID, okID := claims["user_id"].(float64)
		role, okRole := claims["role"].(string)
		exp, okExp := claims["exp"].(float64)

		if okID && okRole && okExp {
			fmt.Printf("UserID: %d, Role: %s, Exp: %f\n", int(userID), role, exp)
		} else {
			fmt.Println("Invalid claims.")
		}

		c.Set("currentUser", user.User{ID: int(userID), Role: role})
		c.Next()
	}
}
