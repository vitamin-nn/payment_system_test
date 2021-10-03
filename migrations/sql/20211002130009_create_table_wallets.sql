-- +goose Up
CREATE TABLE wallets (
    id CHAR(36) NOT NULL PRIMARY KEY,
    currency_code CHAR(3) NOT NULL,
    amount BIGINT NOT NULL,
    user_id CHAR(36) NOT NULL
);
-- +goose Down
DROP TABLE IF EXISTS wallets;
