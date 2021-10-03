package http

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/vitamin-nn/test_payment_system/server/internal/usecase"
)

type HTTP struct {
	useCase *usecase.UseCase
	srv *http.Server
}

func NewHTTP(useCase *usecase.UseCase, wTimeout, rTimeout time.Duration) *HTTP {
	s := new(HTTP)
	s.srv = &http.Server{
		WriteTimeout: wTimeout,
		ReadTimeout:  rTimeout,
		Handler:      s.getRouter(),
	}
	s.useCase = useCase

	return s
}

func (s *HTTP) Run(addr string) error {
	s.srv.Addr = addr

	return s.srv.ListenAndServe()
}

func (s *HTTP) Shutdown(ctx context.Context) error {
	return s.srv.Shutdown(ctx)
}

func (s *HTTP) getRouter() http.Handler {
	router := gin.Default()

	router.POST("/user", s.register)
	router.POST("/rate_update", s.rateUpdate)
	router.POST("/deposit", s.deposit)
	router.POST("/transfer", s.transfer)

	return router
}
