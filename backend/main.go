package main

import (
	"log"
	"os"

	"schoolapp/backend/config"
	"schoolapp/backend/routes"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, relying on system environment variables")
	}

	config.ConnectDB()

	r := routes.SetupRouter()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server starting on port %s", port)
	err = r.Run(":" + port)
	if err != nil {
		log.Fatal("Server failed to start:", err)
	}
}
