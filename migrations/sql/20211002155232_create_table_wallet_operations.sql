-- +goose Up
CREATE TABLE wallet_operations (
    id BIGINT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    wallet_id CHAR(36) NOT NULL,
    amount BIGINT NOT NULL COMMENT "amount in wallet currency",
    amount_usd BIGINT NOT NULL COMMENT "amount in USD",
    amount_operation BIGINT NOT NULL COMMENT "amount in operation currency",
    currency_code CHAR(3) NOT NULL COMMENT "currency operation",
    create_at datetime NOT NULL
);

-- +goose Down
DROP TABLE IF EXISTS wallet_operations;
