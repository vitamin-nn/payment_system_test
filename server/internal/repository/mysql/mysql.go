package mysql

import (
	"context"
	"database/sql"
	"fmt"

	sqlc "github.com/vitamin-nn/test_payment_system/server/internal/repository/mysql/sqlc_generated"
)

type SQLRepo struct {
	db *sql.DB
	*sqlc.Queries
}

func NewSQLRepo(db *sql.DB) *SQLRepo {
	return &SQLRepo{
		db: db,
		Queries: sqlc.New(db),
	}
}

func (sr *SQLRepo) ExecTx(ctx context.Context, fn func(*sqlc.Queries) error) error {
	tx, err := sr.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	qTx := sr.WithTx(tx)
	err = fn(qTx)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("get transaction error: %v, rolback err: %v", err, rbErr)
		}
		return err
	}

	return tx.Commit()
}
