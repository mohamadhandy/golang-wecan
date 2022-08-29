package auth

import (
	"errors"
	"fmt"
	"os"

	"github.com/golang-jwt/jwt/v4"
)

type Service interface {
	GenerateToken(int) (string, error)
	ValidateToken(string) (bool, int, error)
}

type JwtService struct {
	UserId int
	jwt.RegisteredClaims
}

var SECRET_KEY = []byte(os.Getenv("SECRET_KEY"))

func NewService() *JwtService {
	return &JwtService{}
}

func (s *JwtService) GenerateToken(userId int) (string, error) {
	fmt.Println("TOKEN LOGIN", userId)
	claim := JwtService{}
	claim.UserId = userId

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	fmt.Println("TOKEN", token)
	signedToken, err := token.SignedString(SECRET_KEY)
	if err != nil {
		return signedToken, err
	}
	return signedToken, nil
}

func (s *JwtService) ValidateToken(encodedToken string) (bool, int, error) {
	token, err := jwt.ParseWithClaims(encodedToken, &JwtService{}, func(t *jwt.Token) (interface{}, error) {
		_, ok := t.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("invalid token")
		}
		return []byte(SECRET_KEY), nil
	})
	if err != nil {
		return false, -1, err
	}
	if claim, ok := token.Claims.(*JwtService); ok && token.Valid {
		return true, claim.UserId, nil
	} else {
		return false, -1, err
	}
}
