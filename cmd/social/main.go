package main

import (
	"log"
	"os"

	"github.com/jenkka/course-rest-api/internal/app"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = ":8080"
	}

	server := app.Server{
		Address: port,
	}

	server.Start()
}
