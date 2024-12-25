-- name: GetProducts :many
SELECT id, name, description, price, stock FROM products;

-- name: GetProductById :one
SELECT id, name, description, price, stock FROM products WHERE id = $1;

-- name: CreateProduct :one
INSERT INTO products (name, description, price, stock)
VALUES ($1, $2, $3, $4)
RETURNING id, name, description, price, stock, created_at;

-- name: UpdateProduct :exec
UPDATE products
SET name = $1, description = $2, price = $3, stock = $4, updated_at = NOW()
WHERE id = $5
RETURNING id, name, description, price, stock, updated_at;

-- name: DeleteProduct :exec
DELETE FROM products WHERE id = $1;