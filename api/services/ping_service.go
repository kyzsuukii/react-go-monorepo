package services

import (
	"github.com/gin-gonic/gin"
)

type PingService interface {
	Ping(c *gin.Context)
}

type pingService struct{}

func NewPingService() PingService {
	return &pingService{}
}

func (s *pingService) Ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
