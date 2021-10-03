// Code generated by sqlc. DO NOT EDIT.
// source: currency_rate.sql

package db

import (
	"context"
	"time"
)

const changeCurrencyRate = `-- name: ChangeCurrencyRate :exec
REPLACE INTO currency_rates (currency_code, rate, valid_date) VALUES (?, ?, ?)
`

type ChangeCurrencyRateParams struct {
	CurrencyCode string    `json:"currency_code"`
	Rate         int64     `json:"rate"`
	ValidDate    time.Time `json:"valid_date"`
}

func (q *Queries) ChangeCurrencyRate(ctx context.Context, arg ChangeCurrencyRateParams) error {
	_, err := q.db.ExecContext(ctx, changeCurrencyRate, arg.CurrencyCode, arg.Rate, arg.ValidDate)
	return err
}

const getCurrencyRate = `-- name: GetCurrencyRate :one
SELECT rate FROM currency_rates WHERE currency_code = ? AND DATE(valid_date) = DATE(?)
`

type GetCurrencyRateParams struct {
	CurrencyCode string    `json:"currency_code"`
	DATE         time.Time `json:"DATE"`
}

func (q *Queries) GetCurrencyRate(ctx context.Context, arg GetCurrencyRateParams) (int64, error) {
	row := q.db.QueryRowContext(ctx, getCurrencyRate, arg.CurrencyCode, arg.DATE)
	var rate int64
	err := row.Scan(&rate)
	return rate, err
}
