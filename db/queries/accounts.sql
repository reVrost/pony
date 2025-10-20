-- name: CreateAccount :one
INSERT INTO accounts (
    id, alpaca_account_id, status, currency, cash, portfolio_value, buying_power
) VALUES (
    $1, $2, $3, $4, $5, $6, $7
) RETURNING *;

-- name: GetAccount :one
SELECT * FROM accounts WHERE id = $1;

-- name: GetAccountByAlpacaID :one
SELECT * FROM accounts WHERE alpaca_account_id = $1;

-- name: ListAccounts :many
SELECT * FROM accounts ORDER BY created_at DESC;

-- name: UpdateAccount :one
UPDATE accounts SET
    status = $2,
    cash = $3,
    portfolio_value = $4,
    buying_power = $5,
    updated_at = NOW()
WHERE id = $1
RETURNING *;
