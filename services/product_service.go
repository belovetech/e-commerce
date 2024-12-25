package services

import (
	"context"

	"github.com/belovetech/e-commerce/database/sqlc"
)

type ProductService struct {
	queries *sqlc.Queries
}

func NewProductService(queries *sqlc.Queries) *ProductService {
	return &ProductService{queries: queries}
}

func (s *ProductService) GetProducts(ctx context.Context) ([]sqlc.GetProductsRow, error) {
	products, err := s.queries.GetProducts(ctx)
	if err != nil {
		return nil, err
	}
	return products, err
}
