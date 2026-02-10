// Middleware is a function that runs before and after handler for every matching request.
// Request flow:
// request -> middleware 1 (before c.Next()) -> middleware 2 (before c.Next()) -> handler -> middleware 2 (after) -> middleware 1 (after)
package middleware

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func MyGlobalMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// before handler
		fmt.Println("before global middleware")

		// run next middleware / handler
		ctx.Next()

		// after handler
		fmt.Println("after global middleware")
	}
}

func MyGroupMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// before handler
		fmt.Println("before group middleware")

		// run next middleware / handler
		ctx.Next()

		// after handler
		fmt.Println("after group middleware")
	}
}

func MySingleRouteMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// before handler
		fmt.Println("before route middleware")

		// run next middleware / handler
		ctx.Next()

		// after handler
		fmt.Println("after route middleware")
	}
}

func MyMiddlewareTest() {
	r := gin.Default()

	r.Use(MyGlobalMiddleware())

	api := r.Group("/api")
	api.Use(MyGroupMiddleware())
	{
		api.GET("/hello", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "world",
			})
		})
	}

	r.GET("/profile", MySingleRouteMiddleware(), func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "profile",
		})
	})

	if err := r.Run(); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
