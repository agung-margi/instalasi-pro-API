package jwt

import (
	"errors"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Service interface {
	GenerateToken(userID int, role string, expired time.Duration) (string, error)
	ValidateToken(encodedToken string) (*jwt.Token, error)
}

var JWT_SECRET_KEY = []byte(os.Getenv("JWT_SECRET_KEY"))

func GenerateToken(userID int, role string, expired time.Duration) (string, error) {
	claim := jwt.MapClaims{}
	claim["user_id"] = userID
	claim["role"] = role
	claim["exp"] = time.Now().Add(expired).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	signedToken, err := token.SignedString(JWT_SECRET_KEY)

	if err != nil {
		return signedToken, err
	}
	return signedToken, nil
}

func ValidateToken(encodedToken string) (*jwt.Token, error) {
	token, err := jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)

		if !ok {
			return nil, errors.New("invalid token")
		}
		return []byte(JWT_SECRET_KEY), nil
	})
	if err != nil {
		return token, err
	}
	return token, nil
}
