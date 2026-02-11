package middleware

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func LoggingMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// before handler
		start := time.Now()

		// process request
		c.Next()

		// after handler
		duration := time.Since(start)

		status := c.Writer.Status()
		method := c.Request.Method
		path := c.Request.URL.Path
		raw := c.Request.URL.RawQuery
		clientIP := c.ClientIP()

		if raw != "" {
			path = path + "?" + raw
		}

		log.Printf(
			"[GIN] %3d | %13v | %15s | %-7s %s",
			status,
			duration,
			clientIP,
			method,
			path,
		)
	}
}
