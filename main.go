package main

import (
	"github.com/go-playground/validator"
	"github.com/joho/godotenv"
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
	err := godotenv.Load()
	helper.PanicIfError(err)

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

	transactionRepository := repository.NewTransactionRepository()

	//payment
	paymentService := service.NewPaymentService(db, validate, transactionRepository)
	paymentController := controller.NewPaymentController(paymentService)

	//transaction
	transactionService := service.NewTransactionService(transactionRepository, db, validate, userRepository, paymentService)
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
	router.POST("/api/transaction", transactionController.Create)
	router.POST("/api/transaction/notification", paymentController.GetNotification)

	router.PanicHandler = exception.ErrorHandler

	server := http.Server{
		Addr:    "0.0.0.0:3000",
		Handler: middleware.NewAuthMiddleware(router, jwtService),
	}

	err = server.ListenAndServe()
	helper.PanicIfError(err)
}
