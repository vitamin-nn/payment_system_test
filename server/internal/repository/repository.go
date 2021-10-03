package repository

import (
	"context"

	db "github.com/vitamin-nn/test_payment_system/server/internal/repository/mysql/sqlc_generated"
)

type Repo interface {
	db.Querier
	ExecTx(ctx context.Context, fn func(*db.Queries) error) error
	//TransferTx(ctx context.Context, arg TransferTxParams) error
}

type TransferTxParams struct {
	FromAccountID int64 `json:"from_account_id"`
	ToAccountID   int64 `json:"to_account_id"`
	Amount        int64 `json:"amount"`
}
