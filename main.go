package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"code":     200,
			"language": "go version go1.20.6",
			"message":  "Server running on port 3000",
			"version":  "v3.0",
			"description": "belajar fluxcd image update automation",
		})
	})
	r.Run("0.0.0.0:3000")
}
