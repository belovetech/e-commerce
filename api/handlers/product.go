package handlers

import (
	"net/http"

	"github.com/belovetech/e-commerce/database/sqlc"
	"github.com/belovetech/e-commerce/services"
	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	service *services.ProductService
}

type GetProductResponse struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Stock       int64  `json:"stock"`
	Price       string `json:"price"`
}

func NewProductHandler(queries *sqlc.Queries) *ProductHandler {
	service := services.NewProductService(queries)
	return &ProductHandler{service: service}
}

func (h *ProductHandler) GetProducts(c *gin.Context) {
	products, err := h.service.GetProducts(c)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": ErrServer})
		return
	}

	var response []GetProductResponse

	for _, product := range products {
		response = append(response, GetProductResponse{
			ID:          int64(product.ID),
			Name:        product.Name,
			Description: product.Description.String,
			Stock:       int64(product.Stock),
			Price:       product.Price,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"message":  "Products retrieved successfully",
		"products": response,
	})
}
