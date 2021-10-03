-- name: ChangeCurrencyRate :exec
REPLACE INTO currency_rates (currency_code, rate, valid_date) VALUES (?, ?, ?);

-- name: GetCurrencyRate :one
SELECT rate FROM currency_rates WHERE currency_code = ? AND DATE(valid_date) = DATE(?);