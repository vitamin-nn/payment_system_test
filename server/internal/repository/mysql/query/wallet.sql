-- name: CreateWallet :exec
INSERT INTO wallets (
    id,
    currency_code,
    amount,
    user_id
) VALUES (
  ?, ?, 0, ?
);

-- name: AddUserBalance :exec
UPDATE wallets SET amount =  amount + ? WHERE user_id = ?;

-- name: ReduceUserBalance :exec
UPDATE wallets SET amount =  amount - ? WHERE user_id = ?;

-- name: GetWalletByUser :one
SELECT id, currency_code, amount, user_id FROM wallets WHERE user_id = ?;

-- name: GetWalletByUserForUpdate :one
SELECT id, currency_code, amount, user_id FROM wallets WHERE user_id = ? FOR UPDATE;
