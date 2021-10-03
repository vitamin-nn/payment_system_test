package usecase

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	repo "github.com/vitamin-nn/test_payment_system/server/internal/repository"
	db "github.com/vitamin-nn/test_payment_system/server/internal/repository/mysql/sqlc_generated"
)

type UseCase struct {
	dbRepo repo.Repo
}

func NewUseCase(dbRepo repo.Repo) *UseCase {
	return &UseCase{
		dbRepo: dbRepo,
	}
}

func getCurrencyRate(ctx context.Context, q *db.Queries, currencyCode string) (int64, error) {
	argRate := db.GetCurrencyRateParams{
		CurrencyCode: currencyCode,
		DATE: time.Now(),
	}
	rate, err := q.GetCurrencyRate(ctx, argRate)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return 0, fmt.Errorf("currency code: %s is not set for current date", currencyCode)
		}
		return 0, err
	}

	if rate == 0 {
		return 0, fmt.Errorf("unknown currency code: %s", currencyCode)
	}

	return rate, nil
}