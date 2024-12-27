package handlers

import (
	"database/sql"
	"fmt"
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
	userId := utils.GetUserIdFromContext(c)

	if err := c.ShouldBindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": utils.ErrInvalidRequest.Error()})
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

func (h *OrderHandler) GetUserOrders(c *gin.Context) {
	userId := utils.GetUserIdFromContext(c)
	orders, err := h.service.GetUserOrders(c, int32(userId))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		log.Println("Error: ", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Orders retrieved successfully",
		"orders":  orders,
	})
}

// cancel order
type CancelOrderRequest struct {
	OrderId int `json:"order_id" binding:"required"`
}

func (h *OrderHandler) CancelOrder(c *gin.Context) {
	orderId, err := utils.GetIdFromParam(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err = h.service.CancelOrder(c, int32(orderId))
	if err != nil {
		switch err {
		case utils.ErrOrderAlreadyCancelled:
			c.JSON(http.StatusConflict, gin.H{"message": err.Error()})
			return
		case utils.ErrOrderNotPending:
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		case sql.ErrNoRows:
			c.JSON(http.StatusNotFound, gin.H{"message": fmt.Sprintf("Order with id %d not found", orderId)})
			return
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}

	}

	c.JSON(http.StatusOK, gin.H{"message": "Order cancelled successfully"})
}

type UpdateOrderRequest struct {
	Status  string `json:"status" binding:"required"`
	OrderId int    `json:"order_id" binding:"required"`
}

func (h *OrderHandler) UpdateOrderStatus(c *gin.Context) {
	var req UpdateOrderRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": utils.ErrInvalidRequest.Error()})
		log.Println("Error: ", err)
		return
	}

	err := h.service.UpdateOrderStatus(c, int32(req.OrderId), req.Status)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		log.Println("Error: ", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Order status updated successfully"})
}
