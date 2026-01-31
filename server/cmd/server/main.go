package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using environment variables")
	}

	// Initialize configuration
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Println("Starting Unitus server on port " + port)
	
	// TODO: Initialize database
	// TODO: Initialize cache
	// TODO: Setup router
	// TODO: Start server

	// Placeholder
	log.Println("Server implementation to be completed")
}
