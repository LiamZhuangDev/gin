package routing

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// In Gin, form parameters are values sent in the request body using standard HTML form encoding.
// They usually come with content types like:
// - application/x-www-form-urlencoded
// - multipart/form-data

type LoginRequest struct {
	Username string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required"`
	Remember bool   `form:"remember"`
}

//	curl -X POST http://localhost:8080/login \
//	  -H "Content-Type: application/x-www-form-urlencoded" \
//	  -d "username=admin&password=123456&remember=true"
func FormParams() {
	r := gin.Default()

	r.POST("/login", func(ctx *gin.Context) {
		username := ctx.PostForm("username")
		password := ctx.PostForm("password")
		remember := ctx.DefaultPostForm("remember", "false")

		ctx.JSON(http.StatusOK, gin.H{
			"username": username,
			"password": password,
			"remember": remember,
		})
	})

	r.POST("/login/v2", func(ctx *gin.Context) {
		var req LoginRequest
		if err := ctx.ShouldBind(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err,
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"username": req.Username,
			"password": req.Password,
			"remember": req.Remember,
		})
	})

	if err := r.Run(); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}

type LoginJSONRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Remember bool   `json:"remember"`
}

//	curl -X POST http://localhost:8080/login \
//	  -H "Content-Type: application/json" \
//	  -d '{"username":"admin","password":"123456","remember":true}'
func JSONFormParams() {
	r := gin.Default()

	r.POST("/login", func(ctx *gin.Context) {
		var req LoginJSONRequest

		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err,
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"username": req.Username,
			"password": req.Password,
			"remember": req.Remember,
		})
	})

	if err := r.Run(); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}
