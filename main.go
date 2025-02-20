package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World",
		})
	})

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/hello/:name", func(c *gin.Context) {
		name := c.Param("name")

		c.JSON(200, gin.H{
			"message": "Hello " + name,
		})
	})

	// Serve static files from the current directory
	r.StaticFile("/loaderio-8c36ed10c019f659e3e41ad1b3878dd9.txt", "./loaderio-8c36ed10c019f659e3e41ad1b3878dd9.txt")

	r.Run() // listen and serve on 0.0.0.0:8080
}
