package api

import (
	"database/sql"

	"github.com/belovetech/e-commerce/api/handlers"
	"github.com/belovetech/e-commerce/database/sqlc"
	"github.com/belovetech/e-commerce/middlewares"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine, dbConn *sql.DB) {
	queries := sqlc.New(dbConn)
	public := router.Group("/api")

	// handlers
	pingHandler := handlers.NewPingHandler()
	authHandler := handlers.NewAuthHandler(queries)
	userHandler := handlers.NewUserHandler(queries)
	productHandler := handlers.NewProductHandler(queries)
	orderHandler := handlers.NewOrderHandler(dbConn, queries)

	// public routes
	public.GET("/ping", pingHandler.Ping)
	public.POST("/register", authHandler.RegisterUser)
	public.POST("/login", authHandler.Login)

	// protected routes
	protected := router.Group("/api")
	protected.Use(middlewares.AuthMiddleware())

	// order routes
	protected.POST("/orders", orderHandler.CreateOrder)
	protected.GET("/orders", orderHandler.GetUserOrders)
	protected.DELETE("/orders/:id/cancel", orderHandler.CancelOrder)

	// product routes
	protected.GET("/products", productHandler.GetProducts)

	// admin routes
	admin := protected.Group("/admins")
	admin.Use(middlewares.AdminMiddleware())
	admin.GET("/", userHandler.GetAdmins)
	admin.POST("/update-order-status", orderHandler.UpdateOrderStatus)

}
