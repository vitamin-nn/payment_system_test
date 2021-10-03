package report

import (
	"context"
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"strconv"
	"time"

	repo "github.com/vitamin-nn/test_payment_system/server/internal/repository"
	db "github.com/vitamin-nn/test_payment_system/server/internal/repository/mysql/sqlc_generated"
)

const limit = 100

func GenerateCSVReport(ctx context.Context, dbRepo repo.Repo, userID string, beginTime, endTime time.Time, saveIn io.Writer) error {
	writer := csv.NewWriter(saveIn)
	defer writer.Flush()
	writeHeader(writer)

	err := dbRepo.ExecTx(ctx, func(q *db.Queries) error {
		w, err := q.GetWalletByUser(ctx, userID)
		if err != nil {
			return err
		}

		argOperation := db.GetWalletOperationListParams{
			WalletID: w.ID,
			CreateAt: beginTime,
			CreateAt_2: endTime,
			Limit: limit,
			Offset: 0,
		}

		for {
			rowList, err := q.GetWalletOperationList(ctx, argOperation)
			if err != nil {
				return err
			}
		
			for _, row := range rowList {
				err = writeRow(writer, row)
				if err != nil {
					return err
				}
			}
			writer.Flush()

			if len(rowList) < limit {
				break
			}
			argOperation.Offset = argOperation.Offset + limit
		}

		argSumOperation := db.GetSumWalletOperationParams{
			WalletID: w.ID,
			CreateAt: beginTime,
			CreateAt_2: endTime,
		}
		rowSum, err := q.GetSumWalletOperation(ctx, argSumOperation)
		if err != nil {
			return err
		}

		err = writeSum(writer, rowSum)
		if err != nil {
			return err
		}
		writer.Flush()

		return nil
	})

	if err != nil {
		return err
	}

	return nil
}

func writeHeader(w *csv.Writer) error {
	return write(
		w,
		[]string{
			"ID",
			"Amount",
			"Amount USD",
			"Amount operation",
			"Currency code",
			"Create date",
		})
}

func writeRow(w *csv.Writer, row db.WalletOperation) error {
	rowToWrite := []string{
		fmt.Sprintf("%d", row.ID),
		fmt.Sprintf("%v", float64(row.Amount)/100),
		fmt.Sprintf("%v", float64(row.AmountUsd)/100),
		fmt.Sprintf("%v", float64(row.AmountOperation)/100),
		row.CurrencyCode,
		row.CreateAt.String(),
	}
	return write(w, rowToWrite);
}

func writeSum(w *csv.Writer, row db.GetSumWalletOperationRow) error {
	sum, err := convertSumAmount(row.SumAmount)
	if err != nil {
		return err
	}

	sumUsd, err := convertSumAmount(row.SumAmountUsd)
	if err != nil {
		return err
	}

	rowToWrite := []string{
		"Total",
		fmt.Sprintf("%v", float64(sum)/100),
		fmt.Sprintf("%v", float64(sumUsd)/100),
		"",
		"",
		"",
	}

	return write(w, rowToWrite);
}

func write(w *csv.Writer, data []string) error {
	return w.Write(data);
}

func convertSumAmount(sum interface{}) (int64, error) {
	sumSlice, ok := sum.([]uint8)
	if !ok {
		return 0, errors.New("unknown amount format")
	}
	s := string(sumSlice)
	sumInt, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return 0, err
	}

	return sumInt, nil
}