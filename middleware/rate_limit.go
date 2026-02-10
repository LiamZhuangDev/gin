package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

func RateLimitMiddleware() gin.HandlerFunc {
	// allows requests up to rate 10 and permits bursts of at most 20 tokens
	limiter := rate.NewLimiter(10, 20)

	return func(ctx *gin.Context) {
		if !limiter.Allow() {
			ctx.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{
				"error": "Too many requests",
			})
			return
		}
		ctx.Next()
	}
}
