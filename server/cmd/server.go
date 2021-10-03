package cmd

import (
	"context"
	"os"
	"os/signal"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/vitamin-nn/test_payment_system/server/internal/config"

	"github.com/vitamin-nn/test_payment_system/server/internal/http"
	"github.com/vitamin-nn/test_payment_system/server/internal/repository/mysql"
	"github.com/vitamin-nn/test_payment_system/server/internal/usecase"
)

func serverCmd(cfg *config.Config) *cobra.Command {
	return &cobra.Command{
		Use:   "server",
		Short: "Starts payment system server",
		Run: func(cmd *cobra.Command, args []string) {
			log.WithFields(cfg.Fields()).Info("starting payment server")
			dbConn, err := connDB(context.Background(), cfg.MySQL.GetDSN())
			if err != nil {
				log.Fatalf("mysql master connect error: %v", err)
			}

			sqlRepo := mysql.NewSQLRepo(dbConn)
			useCase := usecase.NewUseCase(sqlRepo)
			httpSrv := http.NewHTTP(useCase, cfg.HTTPServer.WriteTimeout, cfg.HTTPServer.ReadTimeout)

			go func() {
				log.Info("starting HTTP server")
				if err := httpSrv.Run(cfg.HTTPServer.GetAddr()); err != nil {
					log.Fatal(err)
				}
			}()

			interruptCh := make(chan os.Signal, 1)
			signal.Notify(interruptCh, os.Interrupt)
			log.Infof("graceful shutdown: %v", <-interruptCh)
			ctx, finish := context.WithTimeout(context.Background(), 5*time.Second)
			defer finish()
			if err := httpSrv.Shutdown(ctx); err != nil {
				log.Error("error while shutdown")
			}

			err = dbConn.Close()
			if err != nil {
				log.Fatalf("mysql close connect error: %v", err)
			}
		},
	}
}
