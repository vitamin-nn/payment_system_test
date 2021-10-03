package usecase

import (
	"context"
	"errors"
	"time"

	"github.com/vitamin-nn/test_payment_system/server/internal/helper"
	db "github.com/vitamin-nn/test_payment_system/server/internal/repository/mysql/sqlc_generated"
)

// - get walets
// - check funds
// - reduce FROM
// - add TO
// - create WO FROM
// - create WO TO

func (uc *UseCase) Transfer(ctx context.Context, userIDFrom, userIDTo, currencyCode string, amount int64) error {
	err := uc.dbRepo.ExecTx(ctx, func(q *db.Queries) error {
		wFrom, err := q.GetWalletByUserForUpdate(ctx, userIDFrom)
		if err != nil {
			return err
		}
		if wFrom.Amount < amount {
			return errors.New("insufficient funds")
		}

		wTo, err := q.GetWalletByUser(ctx, userIDTo)
		if err != nil {
			return err
		}

		rate, err := getCurrencyRate(ctx, q, currencyCode)
		if err != nil {
			return err
		}

		chargedAmount := amount
		if wFrom.CurrencyCode != currencyCode {
			rateFrom, err := getCurrencyRate(ctx, q, wFrom.CurrencyCode)
			if err != nil {
				return err
			}

			chargedAmount = helper.GetConvertedCurrency(rate, rateFrom, amount)
		}

		argReduce := db.ReduceUserBalanceParams{
			Amount: chargedAmount,
			UserID: wFrom.UserID,
		}
		err = q.ReduceUserBalance(ctx, argReduce)
		if err != nil {
			return err
		}

		amountUSD := helper.GetConvertedToUSDAmount(rate, amount)
		argWOFrom := db.CreateWalletOperationParams{
			WalletID: wFrom.ID,
			Amount: -chargedAmount,
			AmountUsd: amountUSD,
			AmountOperation: amount,
			CurrencyCode: currencyCode,
			CreateAt: time.Now(),
		}
		err = q.CreateWalletOperation(ctx, argWOFrom)
		if err != nil {
			return err
		}

		creditedAmount := amount
		if wTo.CurrencyCode != currencyCode {
			rateTo, err := getCurrencyRate(ctx, q, wTo.CurrencyCode)
			if err != nil {
				return err
			}

			creditedAmount = helper.GetConvertedCurrency(rate, rateTo, amount)
		}

		argCredit := db.AddUserBalanceParams{
			Amount: creditedAmount,
			UserID: wTo.UserID,
		}
		err = q.AddUserBalance(ctx, argCredit)
		if err != nil {
			return err
		}

		argWOTo := db.CreateWalletOperationParams{
			WalletID: wTo.ID,
			Amount: creditedAmount,
			AmountUsd: amountUSD,
			AmountOperation: amount,
			CurrencyCode: currencyCode,
			CreateAt: time.Now(),
		}
		err = q.CreateWalletOperation(ctx, argWOTo)
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
