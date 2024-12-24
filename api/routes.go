package api

import (
	"github.com/belovetech/e-commerce/api/handlers"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {

	public := router.Group("/api")

	// Ping route
	pingHandler := handlers.NewPingHandler()
	public.GET("/ping", pingHandler.Ping)

}
