package main

import (
	"github.com/acsauk/go-open-banking/internal/nordigenAPI"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	// load .env file
	err := godotenv.Load()

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	accessBearerToken := nordigenAPI.GetBearerAccessToken()
	log.Println(accessBearerToken)

	nordigenAPI.GetAvailableBanks(accessBearerToken)
}
