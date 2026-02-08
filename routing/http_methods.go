package routing

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HttpMethods() {
	r := gin.Default()

	r.GET("/users", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "get_users",
		})
	})

	r.POST("/users", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "create_user",
		})
	})

	r.PUT("/users/:id", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "update_user",
		})
	})

	r.DELETE("/users/:id", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "delete_user",
		})
	})

	// unlike "PUT" that replaces the whole thing,
	// "PATCH" modifies only the fields in the request.
	// For example, the request below updates only the email and
	// leaves everything else unchanged:
	//
	// PATCH /users/42
	// Content-Type: application/json
	// {
	//   "email": "new@email.com"
	// }
	r.PATCH("/users/:id", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "patch_user",
		})
	})

	r.Any("/orders", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "handle_any",
		})
	})

	r.Handle("GET", "/products", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "get _products",
		})
	})

	if err := r.Run(); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}
