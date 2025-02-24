package jwt

import (
	"errors"
	"fmt"
	"instalasi-pro/configs"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Service interface {
	GenerateToken(userID int, role string, expired time.Duration) (string, error)
	ValidateToken(encodedToken string) (*jwt.Token, error)
}

func GenerateToken(userID int, role string, expired time.Duration) (string, error) {
	var jwtKey = []byte(configs.AppConfig.JWTSecretKey)
	claim := jwt.MapClaims{}
	claim["user_id"] = userID
	claim["role"] = role
	claim["exp"] = time.Now().Add(expired).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	signedToken, err := token.SignedString(jwtKey)
	fmt.Println("JWT_SECRET_KEY:", jwtKey)
	fmt.Println("Signed Token:", signedToken)

	if err != nil {
		return signedToken, err
	}
	return signedToken, nil
}

func ValidateToken(encodedToken string) (*jwt.Token, error) {
	var jwtKey = []byte(configs.AppConfig.JWTSecretKey)
	token, err := jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)

		if !ok {
			return nil, errors.New("invalid token")
		}
		return []byte(jwtKey), nil
	})
	if err != nil {
		return token, err
	}
	return token, nil
}
