// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: products.sql

package sqlc

import (
	"context"
	"database/sql"
	"time"
)

const createProduct = `-- name: CreateProduct :one
INSERT INTO products (name, description, price, stock, created_by)
VALUES ($1, $2, $3, $4, $5)
RETURNING id, name, description, price, stock, created_at, created_by
`

type CreateProductParams struct {
	Name        string         `json:"name"`
	Description sql.NullString `json:"description"`
	Price       string         `json:"price"`
	Stock       int32          `json:"stock"`
	CreatedBy   int32          `json:"created_by"`
}

type CreateProductRow struct {
	ID          int32          `json:"id"`
	Name        string         `json:"name"`
	Description sql.NullString `json:"description"`
	Price       string         `json:"price"`
	Stock       int32          `json:"stock"`
	CreatedAt   time.Time      `json:"created_at"`
	CreatedBy   int32          `json:"created_by"`
}

func (q *Queries) CreateProduct(ctx context.Context, arg CreateProductParams) (CreateProductRow, error) {
	row := q.db.QueryRowContext(ctx, createProduct,
		arg.Name,
		arg.Description,
		arg.Price,
		arg.Stock,
		arg.CreatedBy,
	)
	var i CreateProductRow
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Description,
		&i.Price,
		&i.Stock,
		&i.CreatedAt,
		&i.CreatedBy,
	)
	return i, err
}

const deleteProduct = `-- name: DeleteProduct :exec
DELETE FROM products WHERE id = $1
`

func (q *Queries) DeleteProduct(ctx context.Context, id int32) error {
	_, err := q.db.ExecContext(ctx, deleteProduct, id)
	return err
}

const getProductById = `-- name: GetProductById :one
SELECT id, name, description, price, stock, is_available FROM products WHERE id = $1
`

type GetProductByIdRow struct {
	ID          int32          `json:"id"`
	Name        string         `json:"name"`
	Description sql.NullString `json:"description"`
	Price       string         `json:"price"`
	Stock       int32          `json:"stock"`
	IsAvailable sql.NullBool   `json:"is_available"`
}

func (q *Queries) GetProductById(ctx context.Context, id int32) (GetProductByIdRow, error) {
	row := q.db.QueryRowContext(ctx, getProductById, id)
	var i GetProductByIdRow
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Description,
		&i.Price,
		&i.Stock,
		&i.IsAvailable,
	)
	return i, err
}

const getProducts = `-- name: GetProducts :many
SELECT id, name, description, price, stock, is_available FROM products
`

type GetProductsRow struct {
	ID          int32          `json:"id"`
	Name        string         `json:"name"`
	Description sql.NullString `json:"description"`
	Price       string         `json:"price"`
	Stock       int32          `json:"stock"`
	IsAvailable sql.NullBool   `json:"is_available"`
}

func (q *Queries) GetProducts(ctx context.Context) ([]GetProductsRow, error) {
	rows, err := q.db.QueryContext(ctx, getProducts)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetProductsRow{}
	for rows.Next() {
		var i GetProductsRow
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Description,
			&i.Price,
			&i.Stock,
			&i.IsAvailable,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateProduct = `-- name: UpdateProduct :one
UPDATE products
SET name = $1, description = $2, price = $3, stock = $4, updated_at = NOW()
WHERE id = $5
RETURNING id, name, description, price, stock, updated_at
`

type UpdateProductParams struct {
	Name        string         `json:"name"`
	Description sql.NullString `json:"description"`
	Price       string         `json:"price"`
	Stock       int32          `json:"stock"`
	ID          int32          `json:"id"`
}

type UpdateProductRow struct {
	ID          int32          `json:"id"`
	Name        string         `json:"name"`
	Description sql.NullString `json:"description"`
	Price       string         `json:"price"`
	Stock       int32          `json:"stock"`
	UpdatedAt   time.Time      `json:"updated_at"`
}

func (q *Queries) UpdateProduct(ctx context.Context, arg UpdateProductParams) (UpdateProductRow, error) {
	row := q.db.QueryRowContext(ctx, updateProduct,
		arg.Name,
		arg.Description,
		arg.Price,
		arg.Stock,
		arg.ID,
	)
	var i UpdateProductRow
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Description,
		&i.Price,
		&i.Stock,
		&i.UpdatedAt,
	)
	return i, err
}

const updateProductStock = `-- name: UpdateProductStock :exec
UPDATE products
SET stock = $1, is_available = $2, updated_at = NOW()
WHERE id = $3
`

type UpdateProductStockParams struct {
	Stock       int32        `json:"stock"`
	IsAvailable sql.NullBool `json:"is_available"`
	ID          int32        `json:"id"`
}

func (q *Queries) UpdateProductStock(ctx context.Context, arg UpdateProductStockParams) error {
	_, err := q.db.ExecContext(ctx, updateProductStock, arg.Stock, arg.IsAvailable, arg.ID)
	return err
}
