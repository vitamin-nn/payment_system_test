-- name: CreateWalletOperation :exec
INSERT INTO wallet_operations (
    wallet_id,
    amount,
    amount_usd,
    amount_operation,
    currency_code,
    create_at
) VALUES (
  ?, ?, ?, ?, ?, ?
);

-- name: GetWalletOperationList :many
SELECT id, wallet_id, amount, amount_usd, amount_operation, currency_code, create_at
FROM wallet_operations
WHERE wallet_id = ?
AND create_at > ?
AND create_at < ?
ORDER BY id
LIMIT ? OFFSET ?;

-- name: GetSumWalletOperation :one
SELECT sum(amount) as sum_amount, sum(amount_usd) as sum_amount_usd
FROM wallet_operations
WHERE wallet_id = ?
AND create_at > ?
AND create_at < ?;
