package handlers

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/belovetech/e-commerce/database/sqlc"
	"github.com/belovetech/e-commerce/services"
	"github.com/belovetech/e-commerce/utils"
	"github.com/gin-gonic/gin"
)

type OrderHandler struct {
	service *services.OrderService
}

func NewOrderHandler(db *sql.DB, queries *sqlc.Queries) *OrderHandler {
	service := services.NewOrderService(db, queries)
	return &OrderHandler{service: service}
}

func (h *OrderHandler) CreateOrder(c *gin.Context) {
	var params services.OrderRequest
	var userId int

	user, _ := c.Get("currentUser")
	userId = user.(*utils.User).ID

	if err := c.ShouldBindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		log.Println("Error: ", err)
		return
	}

	order, err := h.service.CreateOrder(c, int32(userId), params.Products)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		log.Println("Error: ", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Order created successfully",
		"order":   order,
	})
}
