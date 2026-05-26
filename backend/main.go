package main

import (
	"log"
	"os"

	"schoolapp/backend/config"
	"schoolapp/backend/routes"
	"schoolapp/backend/services"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, relying on system environment variables")
	}

	config.ConnectDB()

	socketServer := services.InitSocketServer()
	go func() {
		if err := socketServer.Serve(); err != nil {
			log.Fatalf("Socket.IO listen error: %s\n", err)
		}
	}()
	defer socketServer.Close()

	r := routes.SetupRouter(socketServer)

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
