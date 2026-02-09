package routing

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Query parameters are values that come after ? in the URL.
// See example below, note that must use quotes or escape "&".
// curl "http://localhost:8080/search?keyword=golang&page=1&size=10"
func QueryParams() {
	r := gin.Default()

	r.GET("/search", func(ctx *gin.Context) {
		keyword := ctx.Query("keyword")
		page := ctx.DefaultQuery("page", "1")
		size := ctx.Query("size")

		ctx.JSON(http.StatusOK, gin.H{
			"keyword": keyword,
			"page":    page,
			"size":    size,
		})
	})

	if err := r.Run(); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}
