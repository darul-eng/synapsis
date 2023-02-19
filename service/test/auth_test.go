package test

import (
	"context"
	"fmt"
	"github.com/go-playground/validator"
	_ "github.com/lib/pq"
	"tes-synapsis/app"
	user "tes-synapsis/model/api/auth"
	"tes-synapsis/repository"
	"tes-synapsis/service"
	"testing"
)

func TestRegister(t *testing.T) {
	db := app.NewDB()
	validate := validator.New()

	userRepository := repository.NewUserRepository()
	jwtService := service.NewJwtService()
	userService := service.NewAuthService(userRepository, db, validate, jwtService)

	user := user.RegisterRequest{
		Username: "Darul Ikhsan",
		Email:    "darulikhssan@gmail.com",
		Password: "darulikhsan1234",
	}
	response := userService.Register(context.Background(), user)

	fmt.Println(response)
}

func TestLogin(t *testing.T) {
	db := app.NewDB()
	validate := validator.New()

	userRepository := repository.NewUserRepository()
	jwtService := service.NewJwtService()
	userService := service.NewAuthService(userRepository, db, validate, jwtService)

	login := user.LoginRequest{
		Username: "Darul Ikhsan",
		Password: "darulikhan1234",
	}

	response := userService.Login(context.Background(), login)

	fmt.Println(response)
}
