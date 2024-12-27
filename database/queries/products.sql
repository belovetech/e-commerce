-- name: GetProducts :many
SELECT id, name, description, price, stock, is_available FROM products;

-- name: GetProductById :one
SELECT id, name, description, price, stock, is_available FROM products WHERE id = $1;

-- name: CreateProduct :one
INSERT INTO products (name, description, price, stock, created_by)
VALUES ($1, $2, $3, $4, $5)
RETURNING id, name, description, price, stock, created_at, updated_at, created_by;

-- name: UpdateProduct :one
UPDATE products
SET name = $1, description = $2, price = $3, stock = $4, updated_at = NOW()
WHERE id = $5
RETURNING id, name, description, price, stock, updated_at;

-- name: UpdateProductStock :exec
UPDATE products
SET stock = $1, is_available = $2, updated_at = NOW()
WHERE id = $3;

-- name: DeleteProduct :exec
DELETE FROM products WHERE id = $1;
