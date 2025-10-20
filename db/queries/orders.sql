-- name: CreateOrder :one
INSERT INTO orders (
    id, alpaca_order_id, account_id, symbol, side, order_type, qty,
    limit_price, stop_price, time_in_force, status, submitted_at
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12
) RETURNING *;

-- name: GetOrder :one
SELECT * FROM orders WHERE id = $1;

-- name: GetOrderByAlpacaID :one
SELECT * FROM orders WHERE alpaca_order_id = $1;

-- name: ListOrders :many
SELECT * FROM orders
WHERE account_id = $1
ORDER BY created_at DESC
LIMIT $2 OFFSET $3;

-- name: ListOrdersByStatus :many
SELECT * FROM orders
WHERE account_id = $1 AND status = $2
ORDER BY created_at DESC;

-- name: UpdateOrder :one
UPDATE orders SET
    status = $2,
    filled_qty = $3,
    filled_avg_price = $4,
    filled_at = $5,
    canceled_at = $6,
    updated_at = NOW()
WHERE id = $1
RETURNING *;
