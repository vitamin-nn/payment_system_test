package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type Deposit struct {
	UserID string `json:"user_id" binding:"required"`
	Amount int64 `json:"amount" binding:"required"`
}

func (s *HTTP) deposit(c *gin.Context) {
	var req Deposit
	var err error

	err = c.BindJSON(&req)
	if err != nil {
		log.Errorf("bind error: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "bad request"})
		c.Abort()

		return
	}

	err = s.useCase.Deposit(c, req.UserID, req.Amount)
	if err != nil {
		log.Errorf("create deposit db error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal error"})
		c.Abort()

		return
	}

	c.JSON(http.StatusOK, gin.H{"result": "ok"})
}
