package cmd

import (
	"context"
	"io"
	"os"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/spf13/cobra"
	"github.com/vitamin-nn/test_payment_system/server/internal/config"
	"github.com/vitamin-nn/test_payment_system/server/internal/report"
	"github.com/vitamin-nn/test_payment_system/server/internal/repository/mysql"
)

func reportCmd(cfg *config.Config) *cobra.Command {
	var userID, beginTime, endTime, filename string
	command := &cobra.Command{
		Use:   "report",
		Short: "Starts report generating",
		Run: func(cmd *cobra.Command, args []string) {
			if userID == "" {
				log.Fatalf("unknown user id")
			}

			if beginTime == "" {
				beginTime = "1970-01-01T00:00:00Z"
			}
			t1, err := time.Parse(time.RFC3339, beginTime)
			if err != nil {
				log.Fatalf("can not parse begin_time: %v", err)
			}

			var t2 time.Time
			if endTime != "" {
				t2, err = time.Parse(time.RFC3339, endTime)
				if err != nil {
					log.Fatalf("can not parse end_time: %v", err)
				}
			}

			if t2.IsZero() {
				t2 = time.Now()
			}

			dbConn, err := connDB(context.Background(), cfg.MySQL.GetDSN())
			if err != nil {
				log.Fatalf("mysql master connect error: %v", err)
			}

			sqlRepo := mysql.NewSQLRepo(dbConn)

			var f io.WriteCloser
			if filename == "" {
				f = os.Stdout
			} else {
				f, err = os.Create(filename)
				if err != nil {
					log.Fatalf("open file error: %v", err)
				}
				defer f.Close()
			}

			err = report.GenerateCSVReport(context.Background(), sqlRepo, userID, t1, t2, f)
			if err != nil {
				log.Fatalf("csv report generation error: %v", err)
			}
		},
	}

	command.Flags().StringVar(&userID, "user_id", "", "user_id=...")
	err := command.MarkFlagRequired("user_id")
	if err != nil {
		log.Fatalf("user_id required error: %v", err)
	}
	command.Flags().StringVar(&filename, "filename", "", "filename=...")
	command.Flags().StringVar(&beginTime, "begin_time", "1970-01-01T00:00:00Z", "RFC3339: 2006-01-02T15:04:05Z07:00")
	command.Flags().StringVar(&endTime, "end_time", "", "RFC3339: 2006-01-02T15:04:05Z07:00")

	return command
}