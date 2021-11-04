package main

import (
	"fmt"
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
	fmt.Println(accessBearerToken)

	banks := nordigenAPI.GetAvailableBanks(accessBearerToken)
	fmt.Printf("%+v", banks)

	// Build link with monzo - step 4 of quickstart (see if we can just get requisition id from response body)
}
