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
	"tes-synapsis/middleware"
	"tes-synapsis/repository"
	"tes-synapsis/service"
)

func main() {
	db := app.NewDB()
	validate := validator.New()

	//Auth
	userRepository := repository.NewUserRepository()
	jwtService := service.NewJwtService()
	authService := service.NewAuthService(userRepository, db, validate, jwtService)
	authController := controller.NewAuthController(authService)

	//category
	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)

	//product
	productRepository := repository.NewProductRepository()
	productService := service.NewProductService(productRepository, db, validate)
	productController := controller.NewProductController(productService)

	router := httprouter.New()

	router.POST("/api/register", authController.Register)
	router.POST("/api/login", authController.Login)
	router.POST("/api/category", categoryController.Create)
	router.GET("/api/category/:categoryId", categoryController.FindById)
	router.POST("/api/product", productController.Create)
	router.GET("/api/product/:categoryId", productController.FindByCategoryId)
	//router.GET("/api/product/:productId", productController.FindById)

	router.PanicHandler = exception.ErrorHandler

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: middleware.NewAuthMiddleware(router, jwtService),
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
