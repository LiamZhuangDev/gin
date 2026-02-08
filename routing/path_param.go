package routing

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func PathParams() {
	SinglePathParam()
	MultiPathParams()
	WildcardPathParam()
}

// :param matches ONE segment and stops at the next "/"
func SinglePathParam() {
	r := gin.Default()

	r.GET("/users/:id", func(ctx *gin.Context) {
		id := ctx.Param("id")
		ctx.JSON(http.StatusOK, gin.H{
			"id": id,
		})
	})

	if err := r.Run(); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}

func MultiPathParams() {
	r := gin.Default()

	r.GET("/users/:id/posts/:postId", func(ctx *gin.Context) {
		userId := ctx.Param("id")
		postId := ctx.Param("postId")
		ctx.JSON(http.StatusOK, gin.H{
			"userId": userId,
			"postId": postId,
		})
	})

	if err := r.Run(); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}

// "/*" defines a catch-all path parameter. It matches everything after that point in the URL, including "/"
// real uses:
// - static file servers
// - SPA / React / Vue fallback (Send index.html for anything unknown)
// - Proxy / gateway (Forward the reminder of the path)
func WildcardPathParam() {
	r := gin.Default()

	r.GET("/files/*filepath", func(ctx *gin.Context) {
		filePath := ctx.Param("filepath")
		ctx.JSON(http.StatusOK, gin.H{
			"filepath": filePath,
		})
	})

	if err := r.Run(); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}
