
-- name: GetOrderById :one
SELECT id, user_id, total, status FROM orders WHERE user_id = $1;

-- name: CreateOrder :one
INSERT INTO orders (user_id, total, status)
VALUES ($1, $2, 'Pending')
RETURNING id, user_id, total, status, created_at;

-- name: AddOrderItem :exec
INSERT INTO order_items (order_id, product_id, quantity, price)
VALUES ($1, $2, $3, $4);

-- name: CancelOrder :one
UPDATE orders
SET status = 'Cancelled', updated_at = NOW()
WHERE id = $1 AND status = 'Pending'
RETURNING id, user_id, total, status, updated_at;


-- name: UpdateOrderStatus :exec
UPDATE orders
SET status = $1, updated_at = NOW()
WHERE id = $2;


-- name: UpdateOrderTotal :one
UPDATE orders
SET total = (
    SELECT SUM(price * quantity)
    FROM order_items
    WHERE order_id = $1
), updated_at = NOW()
WHERE id = $1
RETURNING id, total, status, updated_at;
