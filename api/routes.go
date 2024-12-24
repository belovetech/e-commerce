package api

import (
	"database/sql"

	"github.com/belovetech/e-commerce/api/handlers"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine, db *sql.DB) {

	public := router.Group("/api")

	// Ping route
	pingHandler := handlers.NewPingHandler()
	public.GET("/ping", pingHandler.Ping)

}
