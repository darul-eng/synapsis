package service

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type jwtServiceImpl struct {
}

var SECRET_KEY = []byte("synapsis_d4rul")

func NewJwtService() JWTService {
	return &jwtServiceImpl{}
}

func (service *jwtServiceImpl) GenerateToken(username string) (string, error) {
	claims := jwt.MapClaims{}
	claims["username"] = username
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(SECRET_KEY)
	if err != nil {
		return tokenString, err
	}
	return tokenString, nil
}

func (service *jwtServiceImpl) VerifyToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// validate signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid token")
		}
		return SECRET_KEY, nil
	})

	if err != nil {
		return token, err
	}

	return token, nil
}
