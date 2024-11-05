package main

import (
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/kyzsuukii/react-go-monorepo/api"
	"github.com/kyzsuukii/react-go-monorepo/api/middleware"
)

var (
	router api.Router = api.NewRouter()
)

func main() {
	app := gin.Default()

	app.SetTrustedProxies(nil)

	app.Use(static.Serve("/", static.LocalFile("./dist", true)))

	app.Use(middleware.CacheControlMiddleware())

	router.SetupRouter(app)

	app.Run(":8080")
}
