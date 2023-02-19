package test

import (
	"fmt"
	"tes-synapsis/helper"
	"tes-synapsis/service"
	"testing"
)

func TestGenerateToken(t *testing.T) {
	jwtService := service.NewJwtService()
	token, err := jwtService.GenerateToken("drlkhsn")
	helper.PanicIfError(err)

	fmt.Println(token)
}

func TestVerifyToken(t *testing.T) {
	jwtService := service.NewJwtService()
	token, err := jwtService.GenerateToken("drlkhsn")
	helper.PanicIfError(err)

	tokens, err := jwtService.VerifyToken(token)

	fmt.Println(tokens.Valid)
}
