package api

import (
	"database/sql"

	"github.com/belovetech/e-commerce/api/handlers"
	"github.com/belovetech/e-commerce/database/sqlc"
	"github.com/belovetech/e-commerce/middlewares"
	"github.com/gin-gonic/gin"
)

const (
	apiVersion = "/api/v1"
)

type Handler struct {
	PingHandler    *handlers.PingHandler
	AuthHandler    *handlers.AuthHandler
	UserHandler    *handlers.UserHandler
	ProductHandler *handlers.ProductHandler
	OrderHandler   *handlers.OrderHandler
}

func initializeHandlers(dbConn *sql.DB, queries *sqlc.Queries) *Handler {
	return &Handler{
		PingHandler:    handlers.NewPingHandler(),
		AuthHandler:    handlers.NewAuthHandler(queries),
		UserHandler:    handlers.NewUserHandler(queries),
		ProductHandler: handlers.NewProductHandler(queries),
		OrderHandler:   handlers.NewOrderHandler(dbConn, queries),
	}
}

func SetupRoutes(router *gin.Engine, dbConn *sql.DB) {
	queries := sqlc.New(dbConn)
	handlers := initializeHandlers(dbConn, queries)

	setupPingRoutes(router, handlers)
	setupAuthRoutes(router, handlers)
	setupOrderRoutes(router, handlers)
	setupProductRoutes(router, handlers)
	setupAdminRoutes(router, handlers)

}

func setupPingRoutes(router *gin.Engine, handlers *Handler) {
	v1 := router.Group(apiVersion)
	v1.GET("/ping", handlers.PingHandler.Ping)
}

func setupAuthRoutes(router *gin.Engine, handlers *Handler) {
	v1 := router.Group(apiVersion)
	auth := v1.Group("/auth")
	{
		auth.POST("/register", handlers.AuthHandler.RegisterUser)
		auth.POST("/login", handlers.AuthHandler.Login)
	}
}

func setupOrderRoutes(router *gin.Engine, handlers *Handler) {
	v1 := router.Group(apiVersion)
	order := v1.Group("/orders")
	{
		order.Use(middlewares.AuthMiddleware())
		order.POST("", handlers.OrderHandler.CreateOrder)
		order.GET("", handlers.OrderHandler.GetUserOrders)
		order.PATCH("/:id/cancel", handlers.OrderHandler.CancelOrder)
	}
}

func setupProductRoutes(router *gin.Engine, handlers *Handler) {
	v1 := router.Group(apiVersion)
	product := v1.Group("/products")
	{
		product.Use(middlewares.AuthMiddleware())
		product.GET("", handlers.ProductHandler.GetProducts)
	}
}

func setupAdminRoutes(router *gin.Engine, handlers *Handler) {
	v1 := router.Group(apiVersion)
	admin := v1.Group("/admins")
	{
		admin.Use(middlewares.AuthMiddleware())
		admin.Use(middlewares.AdminMiddleware())
		admin.GET("", handlers.UserHandler.GetAdmins)
		admin.PATCH("/orders/:id", handlers.OrderHandler.UpdateOrderStatus)
		admin.POST("/products", handlers.ProductHandler.CreateProduct)
		admin.PUT("/products/:id", handlers.ProductHandler.UpdateProduct)
		admin.DELETE("/products/:id", handlers.ProductHandler.DeleteProduct)
	}
}
