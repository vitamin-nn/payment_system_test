-- name: CreateUser :exec
INSERT INTO users (
    id,
    email,
    user_name,
    city,
    country
) VALUES (
  ?, ?, ?, ?, ?
);

-- name: GetUser :one
SELECT id, email, user_name, city, country FROM users WHERE id = ?;
