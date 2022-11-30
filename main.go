package main

import (
	"github.com/gin-gonic/gin"
	"github.com/verssache/hacktiv8-final4/auth"
	"github.com/verssache/hacktiv8-final4/category"
	"github.com/verssache/hacktiv8-final4/handler"
	"github.com/verssache/hacktiv8-final4/helper"
	"github.com/verssache/hacktiv8-final4/product"
	"github.com/verssache/hacktiv8-final4/transactionhistory"
	"github.com/verssache/hacktiv8-final4/user"
)

func main() {
	cfg := helper.LoadConfig()
	db := helper.InitDB()

	authService := auth.NewService()

	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)
	userHandler := handler.NewUserHandler(userService, authService)

	categoryRepository := category.NewRepository(db)
	categoryService := category.NewService(categoryRepository)
	categoryHandler := handler.NewCategoryHandler(categoryService, authService)

	productRepository := product.NewRepository(db)
	productService := product.NewService(productRepository)
	productHandler := handler.NewProductHandler(categoryService, productService, authService)

	transactionRepository := transactionhistory.NewRepository(db)
	transactionService := transactionhistory.NewService(transactionRepository)
	transactionHandler := handler.NewTransactionHistoryHandler(transactionService, userService, productService, categoryService, authService)

	router := gin.Default()
	api := router.Group("/api/v1")

	// User
	api.POST("/users/register", userHandler.RegisterUser)
	api.POST("/users/login", userHandler.LoginUser)
	api.PATCH("/users/topup", authService.AuthMiddleware(authService, userService), userHandler.TopUpBalance)

	// Category
	api.POST("/categories", authService.AuthMiddleware(authService, userService), categoryHandler.CreateCategory)
	api.GET("/categories", authService.AuthMiddleware(authService, userService), categoryHandler.FindAllCategory)
	api.PATCH("/categories/:id", authService.AuthMiddleware(authService, userService), categoryHandler.UpdateCategory)
	api.DELETE("/categories/:id", authService.AuthMiddleware(authService, userService), categoryHandler.DeleteCategory)

	// Product
	api.POST("/products", authService.AuthMiddleware(authService, userService), productHandler.CreateProduct)
	api.GET("/products", authService.AuthMiddleware(authService, userService), productHandler.FindAllProduct)
	api.PUT("/products/:id", authService.AuthMiddleware(authService, userService), productHandler.UpdateProduct)
	api.DELETE("/products/:id", authService.AuthMiddleware(authService, userService), productHandler.DeleteProduct)

	// Transaction
	api.POST("/transactions", authService.AuthMiddleware(authService, userService), transactionHandler.CreateTransaction)
	api.GET("/transactions/my-transactions", authService.AuthMiddleware(authService, userService), transactionHandler.GetMyTransactionHistory)
	api.GET("/transactions/user-transactions", authService.AuthMiddleware(authService, userService), transactionHandler.GetAllTransactionHistory)

	err := router.Run(":" + cfg.ServerPort)
	if err != nil {
		return
	}
}
