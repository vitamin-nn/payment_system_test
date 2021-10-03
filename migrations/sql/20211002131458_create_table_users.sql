-- +goose Up
CREATE TABLE users (
    id CHAR(36) NOT NULL PRIMARY KEY,
    email VARCHAR(256) NOT NULL,
    user_name VARCHAR(256) NOT NULL,
    city VARCHAR(256) NOT NULL,
    country VARCHAR(64) NOT NULL
);
-- +goose Down
DROP TABLE IF EXISTS users;
