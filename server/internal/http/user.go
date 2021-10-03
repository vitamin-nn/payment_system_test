package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type Register struct {
	Email     string    `json:"email" binding:"required"`
	Name string    `json:"user_name" binding:"required"`
	CurrencyCode string `json:"currency_code" binding:"required"`
	City      string    `json:"city"`
	Country      string    `json:"country"`
}

func (s *HTTP) register(c *gin.Context) {
	var req Register
	var err error

	err = c.BindJSON(&req)
	if err != nil {
		log.Errorf("bind error: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "bad request"})
		c.Abort()

		return
	}

	userID, err := s.useCase.Register(c, req.Email, req.Name, req.City, req.Country, req.CurrencyCode)
	if err != nil {
		log.Errorf("create user db error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal error"})
		c.Abort()

		return
	}

	c.JSON(http.StatusOK, gin.H{"user_id": userID})
}
