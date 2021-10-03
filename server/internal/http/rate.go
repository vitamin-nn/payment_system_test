package http

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type RateUpdate struct {
	CurrencyCode string `json:"currency_code" binding:"required"`
	Rate      int64    `json:"rate" binding:"required"`
	ValidDate time.Time `json:"valid_date" binding:"required"`
}

func (s *HTTP) rateUpdate(c *gin.Context) {
	var req RateUpdate
	var err error

	err = c.BindJSON(&req)
	if err != nil {
		log.Errorf("bind error: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "bad request"})
		c.Abort()

		return
	}

	err = s.useCase.RateUpdate(c, req.CurrencyCode, req.Rate, req.ValidDate)
	if err != nil {
		log.Errorf("update rate db error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal error"})
		c.Abort()

		return
	}

	c.JSON(http.StatusOK, gin.H{"result": "ok"})
}
