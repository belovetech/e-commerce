package seeders

import "github.com/belovetech/e-commerce/database/sqlc"

type Seeder interface {
	Name() string
	Seed(queries *sqlc.Queries) error
}

type ProductRequest struct {
	Name        string
	Description string
	Price       int64
	Stock       int64
}
