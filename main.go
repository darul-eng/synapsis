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

	//cart
	cartRepository := repository.NewCartRepository()
	cartService := service.NewCartService(cartRepository, db, validate)
	cartController := controller.NewCartController(cartService)

	//transaction
	transactionRepository := repository.NewTransactionRepository()
	transactionService := service.NewTransactionService(transactionRepository, db, validate)
	transactionController := controller.NewTransactionController(transactionService)

	router := httprouter.New()

	router.POST("/api/register", authController.Register)
	router.POST("/api/login", authController.Login)
	router.POST("/api/category", categoryController.Create)
	router.GET("/api/category/:categoryId", categoryController.FindById)
	router.POST("/api/product", productController.Create)
	router.GET("/api/product/:categoryId", productController.FindByCategoryId)
	router.POST("/api/cart", cartController.Create)
	router.GET("/api/cart", cartController.FindAll)
	router.DELETE("/api/cart/:cartId", cartController.Delete)
	router.POST("/api/cart/shipment", transactionController.Create)

	router.PanicHandler = exception.ErrorHandler

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: middleware.NewAuthMiddleware(router, jwtService),
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
