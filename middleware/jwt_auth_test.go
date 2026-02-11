package middleware_test

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/LiamZhuangDev/gin/middleware"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func setupRouter() *gin.Engine {
	r := gin.Default()

	authorized := r.Group("/api")
	authorized.Use(middleware.AuthMiddleware())
	{
		authorized.GET("/profile", func(c *gin.Context) {
			c.String(http.StatusOK, "access profile allowed")
		})
	}

	return r
}

// In real cases, JWT token should be given by server after client signed in.
func createTestToken() string {
	claims := jwt.MapClaims{
		"user_id": 123,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, _ := token.SignedString(middleware.JWT_KEY)

	return tokenString
}

func TestValidJWTAuth(t *testing.T) {
	gin.SetMode(gin.TestMode)

	r := setupRouter()

	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/api/profile", nil)
	token := createTestToken()
	req.Header.Set("Authorization", "Bearer "+token)

	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("request should pass, got %d", w.Code)
	}
}

func TestInvalidToken(t *testing.T) {
	gin.SetMode(gin.TestMode)

	r := setupRouter()
	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/api/profile", nil)
	req.Header.Set("Authorization", "Bearer "+"wrong token here")
	r.ServeHTTP(w, req)

	if w.Code != http.StatusUnauthorized {
		t.Fatalf("request should pass, got %d", w.Code)
	}
}

func TestNoToken(t *testing.T) {
	gin.SetMode(gin.TestMode)

	r := setupRouter()
	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/api/profile", nil)

	r.ServeHTTP(w, req)

	if w.Code != http.StatusUnauthorized {
		t.Fatalf("request should pass, got %d", w.Code)
	}
}
