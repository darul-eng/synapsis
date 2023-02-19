package service

import "github.com/dgrijalva/jwt-go"

type JWTService interface {
	GenerateToken(username string) (string, error)
	VerifyToken(token string) (*jwt.Token, error)
}
