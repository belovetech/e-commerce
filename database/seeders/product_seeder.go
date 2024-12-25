package seeders

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/belovetech/e-commerce/database/sqlc"
)

type ProductSeeder struct{}

func getProducts() []ProductRequest {
	products := []ProductRequest{
		{
			Name:        "Wireless Mouse",
			Description: "A sleek and ergonomic wireless mouse with adjustable DPI.",
			Price:       2500,
			Stock:       150,
		},
		{
			Name:        "Mechanical Keyboard",
			Description: "A high-quality mechanical keyboard with customizable RGB lighting.",
			Price:       3000,
			Stock:       75,
		},
		{
			Name:        "Gaming Headset",
			Description: "A surround sound gaming headset with noise-canceling microphone.",
			Price:       5000,
			Stock:       200,
		},
		{
			Name:        "Portable SSD",
			Description: "A 1TB portable SSD with high-speed data transfer.",
			Price:       1000,
			Stock:       50,
		},
		{
			Name:        "Smartphone Stand",
			Description: "A durable and adjustable stand for smartphones and tablets.",
			Price:       1500,
			Stock:       300,
		},
	}
	return products
}

func (s ProductSeeder) Name() string {
	return "ProductSeeder"
}
func (s ProductSeeder) Seed(queries *sqlc.Queries) error {
	products := getProducts()
	for _, product := range products {
		productData :=
			sqlc.CreateProductParams{
				Name:        product.Name,
				Description: sql.NullString{String: product.Description, Valid: true},
				Price:       fmt.Sprintf("%d", product.Price),
				Stock:       int32(product.Stock),
			}
		if _, err := queries.CreateProduct(context.Background(), productData); err != nil {
			log.Printf("Failed to create product '%s': %v", product.Name, err)
		} else {
			log.Printf("Product '%s' created successfully", product.Name)
		}
	}
	return nil
}
