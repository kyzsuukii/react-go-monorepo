package middleware

import (
	"strings"

	"github.com/gin-gonic/gin"
)

func CacheControlMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if strings.HasSuffix(c.Request.URL.Path, ".css") || strings.HasSuffix(c.Request.URL.Path, ".js") {
			c.Header("Cache-Control", "public, max-age=86400")
		}
		c.Next()
	}
}
