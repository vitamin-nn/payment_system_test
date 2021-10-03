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
			var t1, t2 time.Time
			var err error
			if beginTime == "" {
				beginTime = "1970-01-01T00:00:00Z"
			}
			t1, err = time.Parse(time.RFC3339, beginTime)
			if endTime != "" {
				t2, err = time.Parse(time.RFC3339, endTime)
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

			report.GenerateCSVReport(context.Background(), sqlRepo, userID, t1, t2, f)
		},
	}

	command.Flags().StringVar(&userID, "user_id", "", "user_id=...")
	command.MarkFlagRequired("user_id")
	command.Flags().StringVar(&filename, "filename", "", "filename=...")
	command.MarkFlagRequired("filename")
	command.Flags().StringVar(&beginTime, "begin_time", "1970-01-01T00:00:00Z", "RFC3339: 2006-01-02T15:04:05Z07:00")
	command.Flags().StringVar(&endTime, "end_time", "", "RFC3339: 2006-01-02T15:04:05Z07:00")

	return command
}