package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)


func main() {
	r := gin.Default()
	r.GET("/get", func(c *gin.Context) {
		c.String(200, "key: %v", "666")
	})

	r.POST("/post", func(context *gin.Context) {
		context.String(200, "post")
	})

	r.PUT("/put", func(context *gin.Context) {
		context.String(http.StatusOK, "put")
	})

	r.DELETE("/delete", func(context *gin.Context) {
		context.String(200, "delete")
	})
	r.Run(":8000")
}
