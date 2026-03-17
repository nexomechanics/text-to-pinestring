package main

import "github.com/gin-gonic/gin"
import "github.com/gin-contrib/cors"

func handleConvert(c *gin.Context) {
	var blocks []block
	var result = ""

	c.ShouldBindJSON(&blocks)
	result, _ = convert(blocks)
	c.JSON(200, gin.H{"result": result})
}

func main() {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"POST", "OPTIONS"},
		AllowHeaders: []string{"Content-Type"},
	}))

	r.POST("/convert", handleConvert)
	r.Run(":8080")
}
