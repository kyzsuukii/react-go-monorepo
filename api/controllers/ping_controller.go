package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/kyzsuukii/react-go-monorepo/api/services"
)

type PingController interface {
	Ping(c *gin.Context)
}

type pingController struct {
	pingService services.PingService
}

func NewPingController(pingService services.PingService) PingController {
	return &pingController{pingService: pingService}
}

func (c *pingController) Ping(ctx *gin.Context) {
	c.pingService.Ping(ctx)
}
