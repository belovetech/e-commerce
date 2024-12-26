package services

import (
	"database/sql"
	"fmt"

	"github.com/belovetech/e-commerce/database/sqlc"
	"github.com/gin-gonic/gin"
)

type OrderService struct {
	queries *sqlc.Queries
	db      *sql.DB
}
type OrderRequest struct {
	Products []OrderProduct `json:"products"`
}

type OrderProduct struct {
	ProductID int `json:"product_id"`
	Quantity  int `json:"quantity"`
}

func NewOrderService(db *sql.DB, queries *sqlc.Queries) *OrderService {
	return &OrderService{queries: queries, db: db}
}

func (s *OrderService) CreateOrder(ctx *gin.Context, userId int32, params []OrderProduct) (sqlc.CreateOrderRow, error) {

	tx, err := s.db.Begin()
	if err != nil {
		return sqlc.CreateOrderRow{}, err
	}

	defer tx.Rollback()
	qtx := s.queries.WithTx(tx)

	order, err := qtx.CreateOrder(ctx, sqlc.CreateOrderParams{
		UserID: int32(userId),
		Total:  "0",
	})

	if err != nil {
		return sqlc.CreateOrderRow{}, err
	}

	for _, param := range params {
		product, err := qtx.GetProductById(ctx, int32(param.ProductID))

		if err != nil {
			if err == sql.ErrNoRows {
				return sqlc.CreateOrderRow{}, fmt.Errorf("product with id %d not found", param.ProductID)
			}

			return sqlc.CreateOrderRow{}, err
		}

		if !product.IsAvailable.Valid || !product.IsAvailable.Bool {
			return sqlc.CreateOrderRow{}, fmt.Errorf("product with id %d is out of stock", param.ProductID)
		}

		err = qtx.AddOrderItem(ctx, sqlc.AddOrderItemParams{
			OrderID:   order.ID,
			ProductID: product.ID,
			Quantity:  int32(param.Quantity),
			Price:     product.Price,
		})

		if err != nil {
			return sqlc.CreateOrderRow{}, err
		}

		productAvailble := product.IsAvailable.Bool
		if product.Stock == 0 {
			productAvailble = false
		}
		qtx.UpdateProductStock(ctx, sqlc.UpdateProductStockParams{
			ID:          product.ID,
			Stock:       product.Stock - int32(param.Quantity),
			IsAvailable: sql.NullBool{Bool: productAvailble, Valid: true},
		})

	}

	updatedOrder, err := qtx.UpdateOrderTotal(ctx, order.ID)
	if err != nil {
		return sqlc.CreateOrderRow{}, err
	}

	return sqlc.CreateOrderRow{
		ID:     updatedOrder.ID,
		UserID: order.UserID,
		Total:  updatedOrder.Total,
		Status: order.Status,
	}, tx.Commit()
}
