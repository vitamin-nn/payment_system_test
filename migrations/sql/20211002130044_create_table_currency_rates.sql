-- +goose Up
CREATE TABLE currency_rates (
    currency_code CHAR(3) NOT NULL,
    rate BIGINT NOT NULL,
    valid_date date NOT NULL,
    primary key (currency_code, valid_date)
);

-- +goose Down
DROP TABLE IF EXISTS currency_rates;
