package main

import (
	"net/http"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func handleConvert(c *gin.Context) {
	apiKey := os.Getenv("API_KEY")
	if apiKey != "" && c.GetHeader("X-API-Key") != apiKey {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	var blocks []block
	c.ShouldBindJSON(&blocks)
	result, _ := convert(blocks)
	c.JSON(200, gin.H{"result": result})
}

func main() {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"POST", "OPTIONS"},
		AllowHeaders: []string{"Content-Type", "X-API-Key"},
	}))

	r.POST("/convert", handleConvert)
	r.Run(":8080")
}
