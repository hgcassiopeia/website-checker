package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hgcassiopeia/website-checker/backend/router"
)

func main() {
	r := router.NewMyRouter()

	r.GET("/test", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.Run(":3000")
}
