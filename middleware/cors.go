// Browsers enforce the Same-Origin Policy and require explicit CORS headers to permit cross-origin response access.
// So Cross-Origin Resource Sharing (CORS) allows a server to tell the browser which origins are permitted to read its responses.
// Real world scenarios where CORS is required:
// - Seperate frontend (http://app.com) & backend (http://api.com)
// - Microservices / API gateway, UI may call multiple domains
// - Third-party integrations, you may want partner sites to call you API.
// - CDN hosting, frontend on CDN, API elsewhere.
package middleware

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func MyCORSMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		origin := ctx.Request.Header.Get("Origin")
		if origin != "" {
			ctx.Header("Access-Control-Allow-Origin", origin) // The server says "I allow whoever asks". Should use whitelist in production.
			ctx.Header("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS")
			ctx.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")
			ctx.Header("Access-Control-Allow-Credentials", "true")
		}

		// return success for preflight request
		if ctx.Request.Method == "OPTIONS" {
			ctx.AbortWithStatus(http.StatusNoContent)
			return
		}

		ctx.Next()
	}
}

func GinOfficialCORSMiddleware(r *gin.Engine) {
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000"}
	config.AllowCredentials = true

	r.Use(cors.New(config))
}
