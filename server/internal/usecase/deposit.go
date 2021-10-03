package usecase

import (
	"context"
	"log"
	"time"

	"github.com/vitamin-nn/test_payment_system/server/internal/helper"
	db "github.com/vitamin-nn/test_payment_system/server/internal/repository/mysql/sqlc_generated"
)

func (uc *UseCase) Deposit(ctx context.Context, userID string, amount int64) error {
	err := uc.dbRepo.ExecTx(ctx, func(q *db.Queries) error {
		argWallet := db.AddUserBalanceParams{
			Amount: amount,
			UserID: userID,
		}
		err := q.AddUserBalance(ctx, argWallet)
		if err != nil {
			log.Println("add balance err")
			return err
		}

		w, err := q.GetWalletByUser(ctx, userID)
		if err != nil {
			return err
		}

		rate, err := getCurrencyRate(ctx, q, w.CurrencyCode)
		if err != nil {
			return err
		}
		
		argWO := db.CreateWalletOperationParams{
			WalletID: w.ID,
			Amount: amount,
			AmountUsd: helper.GetConvertedToUSDAmount(rate, amount),
			AmountOperation: amount,
			CurrencyCode: w.CurrencyCode,
			CreateAt: time.Now(),
		}

		err = q.CreateWalletOperation(ctx, argWO)
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return err
	}

	return nil
}