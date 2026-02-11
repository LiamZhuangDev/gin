package middleware

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestRecoveryMiddleware(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.Use(RecoveryMiddleware())
	r.GET("/panic", func(ctx *gin.Context) {
		a := 0
		boom := 10 / a
		ctx.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("test recovery middleware, %.2f", float64(boom)),
		})
	})

	req := httptest.NewRequest("GET", "/panic", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusInternalServerError {
		t.Fatalf("request should pass, got %d", w.Code)
	}
}
