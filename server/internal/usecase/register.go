package usecase

import (
	"context"

	"github.com/vitamin-nn/test_payment_system/server/internal/helper"
	db "github.com/vitamin-nn/test_payment_system/server/internal/repository/mysql/sqlc_generated"
)

func (uc *UseCase) Register(ctx context.Context, email, username, city, country, currencyCode string) (string, error) {
	userID := helper.GetUUID()
	argUser := db.CreateUserParams{
		ID: userID,
		Email:       email,
		UserName:       username,
		City:       city,
		Country:          country,
	}

	argWallet := db.CreateWalletParams{
		ID: helper.GetUUID(),
		CurrencyCode: currencyCode,
		UserID: userID,
	}

	err := uc.dbRepo.ExecTx(ctx, func(q *db.Queries) error {
		err := q.CreateUser(ctx, argUser)
		if err != nil {
			return err
		}
	
		err = q.CreateWallet(ctx, argWallet)
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return "", err
	}

	return userID, nil
}