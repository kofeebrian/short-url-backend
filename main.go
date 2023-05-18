package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/kofeebrian/short-url-server/routes"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Printf("Error loading .env file: %v", err)
		log.Printf("Will use environment variables")
	}

	r := routes.SetupRouters()

	r.Run(":8080")
}
