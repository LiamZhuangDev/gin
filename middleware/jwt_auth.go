// 1. JSON Web Token (JWT) is a way to authenticate requests using a signed token instead of sever-stored sessions.
// 2. user logs in once -> server gives a token -> client sends token on every request -> server verifies -> allow or deny
// 3. JWT has 3 parts: header, payload and signature.
// payload is NOT encrypted, only signed. Anyone can read it, but they can't modify it without breaking the signature.
// payload example:
// {
//   "user_id": 123,
//   "role": "admin",
//   "exp": 1735689600
// }

package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var JWT_KEY = []byte("secret")

func Login(c *gin.Context) {
	// after verifying username/password
	tokenString := "see creatTestToken() for an example"
	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := c.GetHeader("Authorization")

		if auth == "" {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		tokenString := strings.TrimPrefix(auth, "Bearer ")

		token, err := jwt.Parse(tokenString, func(t *jwt.Token) (any, error) {
			return JWT_KEY, nil
		})

		if err != nil || !token.Valid {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		c.Next()
	}
}
