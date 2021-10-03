package usecase

import (
	"context"
	"time"

	db "github.com/vitamin-nn/test_payment_system/server/internal/repository/mysql/sqlc_generated"
)

func (uc *UseCase) RateUpdate(ctx context.Context, currencyCode string, rate int64, validDate time.Time) error {
	argChangeRate := db.ChangeCurrencyRateParams{
		CurrencyCode: currencyCode,
		Rate: rate,
		ValidDate: validDate,
	}
	err := uc.dbRepo.ChangeCurrencyRate(ctx, argChangeRate)
	if err != nil {
		return err
	}

	return nil
}