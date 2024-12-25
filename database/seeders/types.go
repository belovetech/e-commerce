package seeders

import (
	"github.com/belovetech/e-commerce/config"
	"github.com/belovetech/e-commerce/database/sqlc"
)

type Seeder interface {
	Name() string
	Seed(queries *sqlc.Queries, cfg *config.Config) error
}

type ProductRequest struct {
	Name        string
	Description string
	Price       int64
	Stock       int64
}
