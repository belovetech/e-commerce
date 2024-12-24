package api

import (
	"database/sql"

	"github.com/belovetech/e-commerce/api/handlers"
	"github.com/belovetech/e-commerce/database/sqlc"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine, dbConn *sql.DB) {
	queries := sqlc.New(dbConn)

	public := router.Group("/api")

	// Ping route
	pingHandler := handlers.NewPingHandler()
	public.GET("/ping", pingHandler.Ping)

	// users route
	userHandler := handlers.NewUserHandler(queries)
	public.POST("/register", userHandler.RegisterUser)

	// protected routes
	protected := router.Group("/api")
	// protected.Use(AuthMiddleware())
	protected.GET("/admins", userHandler.GetAdmins)

}
