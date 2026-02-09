package routing

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Routes grouping means organizing multiple routes under a common path prefix, often sharing middleware as well.
// A group is like a folder: folder path + file path = full path

func GroupBasis() {
	r := gin.Default()

	// Group following endpoints:
	// r.GET("/api/v1/users", getUsers)
	// r.POST("/api/v1/users", createUser)
	// r.GET("/api/v1/orders", getOrders)
	api := r.Group("/api/v1")
	{
		getUsers := func(c *gin.Context) {
			c.JSON(http.StatusOK, []gin.H{
				{"id": 1, "name": "Alice"},
				{"id": 2, "name": "Bob"},
			})
		}

		getOrders := func(c *gin.Context) {
			c.JSON(http.StatusOK, []gin.H{
				{"id": 101, "item": "Keyboard", "price": 99.9},
				{"id": 102, "item": "Mouse", "price": 49.5},
			})
		}

		api.GET("/users", getUsers)
		api.GET("/orders", getOrders)
	}

	if err := r.Run(); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}

func Group4Middleware() {
	r := gin.Default()

	api := r.Group("/api")
	api.Use(AuthMiddleware())
	{
		getProfile := func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"name":    "Alice",
				"address": "123 main st. LA, CA",
			})
		}

		getSettings := func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"WiFi":             "on",
				"Network Provider": "AT&T",
			})
		}

		api.GET("/profile", getProfile)
		api.GET("/settings", getSettings)
	}

	if err := r.Run(); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}

//	curl -H "Authorization: Bearer demo-token" \
//	     http://localhost:8080/api/profile
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := c.GetHeader("Authorization")

		// very fake validation, just for demo
		if auth != "Bearer demo-token" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "unauthorized",
			})
			return
		}

		// continue to the next handler
		c.Next()
	}
}

func NestedGroup() {
	r := gin.Default()

	api := r.Group("/api")
	{
		getUsers := func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"users": "v1",
			})
		}

		getUsersV2 := func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"users": "v2",
			})
		}

		v1 := api.Group("/v1")
		{
			v1.GET("/users", getUsers)
			// other handlers
		}

		v2 := api.Group("/v2")
		{
			v2.GET("/users", getUsersV2)
			// other V2 handlers
		}
	}

	if err := r.Run(); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}
