package api

import (
	"github.com/gin-gonic/gin"
	"github.com/kyzsuukii/react-go-monorepo/api/controllers"
	"github.com/kyzsuukii/react-go-monorepo/api/services"
)

type Router interface {
	SetupRouter(router *gin.Engine)
}

type router struct {
	pingController controllers.PingController
}

func NewRouter() Router {
	pingService := services.NewPingService()
	pingController := controllers.NewPingController(pingService)
	return &router{pingController: pingController}
}

func (r *router) SetupRouter(router *gin.Engine) {
	apiGroup := router.Group("/api")
	{
		apiGroup.GET("/ping", r.pingController.Ping)
	}
}
