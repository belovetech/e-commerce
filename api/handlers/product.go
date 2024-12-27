package handlers

import (
	"database/sql"
	"log"
	"net/http"
	"time"

	"github.com/belovetech/e-commerce/database/sqlc"
	"github.com/belovetech/e-commerce/services"
	"github.com/belovetech/e-commerce/utils"
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

type CreateProductRequest struct {
	Stock       int32  `json:"stock" binding:"required,min=0"`
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"max=500"`
	Price       string `json:"price" binding:"required"`
}

type ProductResponse struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Stock       int64     `json:"stock"`
	Price       string    `json:"price"`
	CreatedBy   int64     `json:"created_by"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (h *ProductHandler) CreateProduct(c *gin.Context) {
	var params CreateProductRequest
	userId := utils.GetUserIdFromContext(c)

	if err := c.ShouldBindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": ErrInvalidRequest, "error": err.Error()})
		log.Printf("Error while binding JSON: %v\n", err)
		return
	}

	product, err := h.service.CreateProduct(c, sqlc.CreateProductParams{
		Name:        params.Name,
		Description: sql.NullString{String: params.Description, Valid: true},
		Stock:       params.Stock,
		Price:       params.Price,
		CreatedBy:   int32(userId),
	})

	if err != nil {
		log.Printf("Error while creating product: %v\n", err)
		if err == utils.ErrProductExists {
			c.JSON(http.StatusConflict, gin.H{"message": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"message": ErrServer})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Product created successfully",
		"product": ProductResponse{
			ID:          int64(product.ID),
			Name:        product.Name,
			Description: product.Description.String,
			Stock:       int64(product.Stock),
			Price:       product.Price,
			CreatedBy:   int64(product.CreatedBy),
			CreatedAt:   product.CreatedAt,
		},
	})
}

type UpdateProductRequest struct {
	Stock       int32  `json:"stock" binding:"min=0,omitempty"`
	Name        string `json:"name" binding:"omitempty"`
	Description string `json:"description" binding:"max=500,omitempty"`
	Price       string `json:"price" binding:"omitempty"`
}

func (h *ProductHandler) UpdateProduct(c *gin.Context) {
	productId, err := utils.GetIdFromParam(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		log.Printf("Error while getting product ID: %v\n", err)
		return
	}

	var params UpdateProductRequest
	if err := c.ShouldBindJSON(&params); err != nil {
		log.Printf("Error while binding JSON: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"message": ErrInvalidRequest})
		return
	}

	product, err := h.service.UpdateProduct(c, sqlc.UpdateProductParams{
		ID:          int32(productId),
		Name:        params.Name,
		Description: sql.NullString{String: params.Description, Valid: true},
		Stock:       params.Stock,
		Price:       params.Price,
	})

	if err != nil {
		log.Printf("Error while updating product: %v\n", err)
		if err == utils.ErrProductNotFound {
			c.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"message": ErrServer})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Product updated successfully",
		"product": ProductResponse{
			ID:          int64(product.ID),
			Name:        product.Name,
			Description: product.Description.String,
			Stock:       int64(product.Stock),
			Price:       product.Price,
			UpdatedAt:   product.UpdatedAt,
		},
	})
}

func (h *ProductHandler) DeleteProduct(c *gin.Context) {
	productId, err := utils.GetIdFromParam(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = h.service.DeleteProduct(c, productId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": ErrServer})
		return
	}

	c.JSON(http.StatusNoContent, gin.H{
		"message": "Product deleted successfully",
	})
}
