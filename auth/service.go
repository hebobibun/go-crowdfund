package auth

import (
	"crowdfund/config"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type Service interface {
	GenerateToken(userID int) (string, error)
}

type jwtService struct {
}

func NewJWT() *jwtService {
	return &jwtService{}
}

func (s *jwtService) GenerateToken(userID int) (string, error) {
	claim := jwt.MapClaims{}
	claim["authorized"] = true
	claim["user_id"] = userID
	claim["exp"] = time.Now().Add(time.Hour * 48).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	signedToken, err := token.SignedString([]byte(config.JWTKEY))
	if err != nil {
		return signedToken, err
	}

	return signedToken, nil
}
