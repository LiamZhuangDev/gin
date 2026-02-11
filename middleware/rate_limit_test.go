package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

func TestRateLimitMiddleware(t *testing.T) {
	gin.SetMode(gin.TestMode)

	r := gin.New()
	// allows requests up to rate 10 and permits bursts of at most 20 tokens
	limiter := rate.NewLimiter(10, 20)
	r.Use(RateLimitMiddleware(limiter))
	r.GET("/ping", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "pong")
	})

	// send 21 requests immediately
	for i := 1; i <= 21; i++ {
		req := httptest.NewRequest("GET", "/ping", nil)
		w := httptest.NewRecorder()

		r.ServeHTTP(w, req)

		if i <= 20 && w.Code != http.StatusOK {
			t.Fatalf("request %d should pass, got %d", i, w.Code)
		}

		if i == 21 && w.Code != http.StatusTooManyRequests {
			t.Fatalf("request %d should be blocked, got %d", i, w.Code)
		}
	}
}
