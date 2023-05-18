package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/kofeebrian/short-url-server/routes"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	r := routes.SetupRouters()

	r.Run(":8080")
}
