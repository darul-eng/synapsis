package main

import (
	"github.com/go-playground/validator"
	"github.com/julienschmidt/httprouter"
	_ "github.com/lib/pq"
	"net/http"
	"tes-synapsis/app"
	"tes-synapsis/controller"
	"tes-synapsis/exception"
	"tes-synapsis/helper"
	"tes-synapsis/repository"
	"tes-synapsis/service"
)

func main() {
	db := app.NewDB()
	validate := validator.New()
	userRepository := repository.NewUserRepository()
	jwtService := service.NewJwtService()
	authService := service.NewAuthService(userRepository, db, validate, jwtService)
	authController := controller.NewAuthController(authService)

	router := httprouter.New()

	router.POST("/api/register", authController.Register)
	router.POST("/api/login", authController.Login)

	router.PanicHandler = exception.ErrorHandler

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: router,
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
