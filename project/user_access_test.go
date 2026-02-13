package project

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/require"
)

func setupRouter() *gin.Engine {
	r := gin.Default()

	authorized := r.Group("/api")
	authorized.Use(JwtAuthMiddleware())
	{
		authorized.GET("/profile", func(ctx *gin.Context) {
			role := ctx.GetString("role")
			if role == "" {
				ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
				return
			}
			if role != "admin" {
				ctx.JSON(http.StatusForbidden, gin.H{"error": "Forbidden"})
				return
			}

			ctx.JSON(200, gin.H{
				"role": role,
			})
		})
	}

	return r
}

var JWT_KEY = []byte("secret")

type MyClaims struct {
	UserID int    `json:"user_id"`
	Role   string `json:"role"`
	jwt.RegisteredClaims
}

func JwtAuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		auth := ctx.GetHeader("Authorization")

		if auth == "" {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		tokenString := strings.TrimPrefix(auth, "Bearer ")
		claims := &MyClaims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (any, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.ErrSignatureInvalid
			}

			return JWT_KEY, nil
		})
		if err != nil || !token.Valid {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		ctx.Set("role", claims.Role)

		ctx.Next()
	}
}

// In real cases, JWT token should be given by server after client signed in.
func createTestToken(role string) string {
	claims := MyClaims{
		UserID: 123,
		Role:   role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	if tokenString, err := token.SignedString(JWT_KEY); err != nil {
		panic(err)
	} else {
		return tokenString
	}
}

func TestAdminAccess(t *testing.T) {
	gin.SetMode(gin.TestMode)

	r := setupRouter()

	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/api/profile", nil)
	token := createTestToken("admin")
	req.Header.Set("Authorization", "Bearer "+token)

	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code)
}
