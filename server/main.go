package main

import (
	"log"
	"server/database"
	"server/routes"
)

func main() {
	// Initialize the database
	database.InitDB()

	// Set up routes
	router := routes.SetupRouter()

	// Start the server
	log.Println("Server running on http://localhost:8080")
	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
