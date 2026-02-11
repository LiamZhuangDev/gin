package middleware_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/LiamZhuangDev/gin/middleware"
	"github.com/gin-gonic/gin"
)

func TestCORS_Preflight(t *testing.T) {
	gin.SetMode(gin.TestMode)

	r := gin.New()
	r.Use(middleware.MyCORSMiddleware())

	// this handler should NOT run for OPTIONS
	r.GET("/ping", func(c *gin.Context) {
		t.Fatal("handler should not be called for preflight")
	})

	req := httptest.NewRequest(http.MethodOptions, "/ping", nil)
	req.Header.Set("Origin", "https://evil.com")

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// should stop at middleware
	if w.Code != http.StatusNoContent {
		t.Fatalf("expected 204, got %d", w.Code)
	}

	// headers should exist
	if w.Header().Get("Access-Control-Allow-Origin") != "https://evil.com" {
		t.Fatal("missing allow-origin header")
	}
}

func TestCORS_NormalRequest(t *testing.T) {
	gin.SetMode(gin.TestMode)

	r := gin.New()
	r.Use(middleware.MyCORSMiddleware())

	handlerCalled := false

	r.GET("/ping", func(c *gin.Context) {
		handlerCalled = true
		c.String(200, "pong")
	})

	req := httptest.NewRequest(http.MethodGet, "/ping", nil)
	req.Header.Set("Origin", "https://example.com")

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", w.Code)
	}

	if !handlerCalled {
		t.Fatal("handler should be called")
	}

	if w.Header().Get("Access-Control-Allow-Origin") != "https://example.com" {
		t.Fatal("header not set")
	}
}
