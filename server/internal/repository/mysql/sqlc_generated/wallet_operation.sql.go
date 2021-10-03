// Code generated by sqlc. DO NOT EDIT.
// source: wallet_operation.sql

package db

import (
	"context"
	"time"
)

const createWalletOperation = `-- name: CreateWalletOperation :exec
INSERT INTO wallet_operations (
    wallet_id,
    amount,
    amount_usd,
    amount_operation,
    currency_code,
    create_at
) VALUES (
  ?, ?, ?, ?, ?, ?
)
`

type CreateWalletOperationParams struct {
	WalletID        string    `json:"wallet_id"`
	Amount          int64     `json:"amount"`
	AmountUsd       int64     `json:"amount_usd"`
	AmountOperation int64     `json:"amount_operation"`
	CurrencyCode    string    `json:"currency_code"`
	CreateAt        time.Time `json:"create_at"`
}

func (q *Queries) CreateWalletOperation(ctx context.Context, arg CreateWalletOperationParams) error {
	_, err := q.db.ExecContext(ctx, createWalletOperation,
		arg.WalletID,
		arg.Amount,
		arg.AmountUsd,
		arg.AmountOperation,
		arg.CurrencyCode,
		arg.CreateAt,
	)
	return err
}

const getSumWalletOperation = `-- name: GetSumWalletOperation :one
SELECT sum(amount) as sum_amount, sum(amount_usd) as sum_amount_usd
FROM wallet_operations
WHERE wallet_id = ?
AND create_at > ?
AND create_at < ?
`

type GetSumWalletOperationParams struct {
	WalletID   string    `json:"wallet_id"`
	CreateAt   time.Time `json:"create_at"`
	CreateAt_2 time.Time `json:"create_at_2"`
}

type GetSumWalletOperationRow struct {
	SumAmount    interface{} `json:"sum_amount"`
	SumAmountUsd interface{} `json:"sum_amount_usd"`
}

func (q *Queries) GetSumWalletOperation(ctx context.Context, arg GetSumWalletOperationParams) (GetSumWalletOperationRow, error) {
	row := q.db.QueryRowContext(ctx, getSumWalletOperation, arg.WalletID, arg.CreateAt, arg.CreateAt_2)
	var i GetSumWalletOperationRow
	err := row.Scan(&i.SumAmount, &i.SumAmountUsd)
	return i, err
}

const getWalletOperationList = `-- name: GetWalletOperationList :many
SELECT id, wallet_id, amount, amount_usd, amount_operation, currency_code, create_at
FROM wallet_operations
WHERE wallet_id = ?
AND create_at > ?
AND create_at < ?
ORDER BY id
LIMIT ? OFFSET ?
`

type GetWalletOperationListParams struct {
	WalletID   string    `json:"wallet_id"`
	CreateAt   time.Time `json:"create_at"`
	CreateAt_2 time.Time `json:"create_at_2"`
	Limit      int32     `json:"limit"`
	Offset     int32     `json:"offset"`
}

func (q *Queries) GetWalletOperationList(ctx context.Context, arg GetWalletOperationListParams) ([]WalletOperation, error) {
	rows, err := q.db.QueryContext(ctx, getWalletOperationList,
		arg.WalletID,
		arg.CreateAt,
		arg.CreateAt_2,
		arg.Limit,
		arg.Offset,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []WalletOperation{}
	for rows.Next() {
		var i WalletOperation
		if err := rows.Scan(
			&i.ID,
			&i.WalletID,
			&i.Amount,
			&i.AmountUsd,
			&i.AmountOperation,
			&i.CurrencyCode,
			&i.CreateAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
