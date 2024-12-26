package services

import (
	"context"
	"log"

	"github.com/belovetech/e-commerce/database/sqlc"
)

type ProductService struct {
	queries *sqlc.Queries
}

func NewProductService(queries *sqlc.Queries) *ProductService {
	return &ProductService{queries: queries}
}

func (s *ProductService) CreateProduct(ctx context.Context, params sqlc.CreateProductParams) (sqlc.CreateProductRow, error) {
	product, err := s.queries.CreateProduct(ctx, params)

	log.Printf("Product: %v", product)
	if err != nil {
		return sqlc.CreateProductRow{}, err
	}
	return product, err
}

func (s *ProductService) GetProducts(ctx context.Context) ([]sqlc.GetProductsRow, error) {
	products, err := s.queries.GetProducts(ctx)
	if err != nil {
		return nil, err
	}
	return products, err
}

func (s *ProductService) GetProductById(ctx context.Context, id int32) (sqlc.GetProductByIdRow, error) {
	product, err := s.queries.GetProductById(ctx, id)
	if err != nil {
		return sqlc.GetProductByIdRow{}, err
	}
	return product, err
}

func (s *ProductService) UpdateProduct(ctx context.Context, params sqlc.UpdateProductParams) (sqlc.UpdateProductRow, error) {
	product, err := s.queries.UpdateProduct(ctx, params)
	if err != nil {
		return sqlc.UpdateProductRow{}, err
	}
	return product, err
}

func (s *ProductService) DeleteProduct(ctx context.Context, id int32) error {
	err := s.queries.DeleteProduct(ctx, id)
	if err != nil {
		return err
	}
	return nil
}
