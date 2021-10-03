package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type Transfer struct {
	UserIDFrom string `json:"user_id_from" binding:"required"`
	UserIDTo string `json:"user_id_to" binding:"required"`
	Amount int64 `json:"amount" binding:"required"`
	CurrencyCode string `json:"currency_code" binding:"required"`
}

func (s *HTTP) transfer(c *gin.Context) {
	var req Transfer
	var err error

	err = c.BindJSON(&req)
	if err != nil {
		log.Errorf("bind error: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "bad request"})
		c.Abort()

		return
	}

	err = s.useCase.Transfer(c, req.UserIDFrom, req.UserIDTo, req.CurrencyCode, req.Amount)
	if err != nil {
		log.Errorf("create transfer db error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal error"})
		c.Abort()

		return
	}

	c.JSON(http.StatusOK, gin.H{"result": "ok"})
}
