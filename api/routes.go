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

	// Ping route
	pingHandler := handlers.NewPingHandler()
	public.GET("/ping", pingHandler.Ping)

	// auth route
	authHandler := handlers.NewAuthHandler(queries)
	public.POST("/register", authHandler.RegisterUser)
	public.POST("/login", authHandler.Login)

	// protected routes
	userHandler := handlers.NewUserHandler(queries)
	protected := router.Group("/api")
	protected.Use(middlewares.AuthMiddleware())

	// products route
	productHandler := handlers.NewProductHandler(queries)
	protected.GET("/products", productHandler.GetProducts)

	// admin routes
	protected.Use(middlewares.AdminMiddleware())
	protected.GET("/admins", userHandler.GetAdmins)

}
