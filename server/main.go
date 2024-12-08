package main

import (
	"log"
	"github.com/gin-gonic/gin"
)

func main() {
	// Create a Gin router
	router := gin.Default()

	// Define the route for GET request at the root
	router.GET("/", func(c *gin.Context) {
		c.String(200, "Hello, World!")
	})

	// Log that the server is running
	log.Println("Server running on http://localhost:8080")

	// Start the Gin server on port 8080
	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
