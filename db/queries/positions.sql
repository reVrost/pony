-- name: CreatePosition :one
INSERT INTO positions (
    account_id, symbol, qty, avg_entry_price, current_price,
    market_value, cost_basis, unrealized_pl, unrealized_plpc
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, $9
) RETURNING *;

-- name: GetPosition :one
SELECT * FROM positions
WHERE account_id = $1 AND symbol = $2;

-- name: ListPositions :many
SELECT * FROM positions
WHERE account_id = $1
ORDER BY market_value DESC;

-- name: UpdatePosition :one
UPDATE positions SET
    qty = $3,
    avg_entry_price = $4,
    current_price = $5,
    market_value = $6,
    cost_basis = $7,
    unrealized_pl = $8,
    unrealized_plpc = $9,
    updated_at = NOW()
WHERE account_id = $1 AND symbol = $2
RETURNING *;

-- name: DeletePosition :exec
DELETE FROM positions
WHERE account_id = $1 AND symbol = $2;
