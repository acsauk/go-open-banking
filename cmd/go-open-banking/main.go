package main

import (
	"fmt"
	"github.com/acsauk/go-open-banking/internal/nordigenAPI"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	// load .env file
	err := godotenv.Load()

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	accessBearerToken := nordigenAPI.GetBearerAccessToken()
	//fmt.Println(accessBearerToken)

	banks := nordigenAPI.GetAvailableBanks(accessBearerToken)
	//fmt.Printf("%+v", banks)

	// Build link with monzo - step 4 of quickstart (see if we can just get requisition id from response body)
	var bankId string

	for _, v := range banks {
		if v.Name == "Monzo Bank Limited" {
			bankId = v.Id
		}
	}

	if bankId == "" {
		log.Fatalf("Could not find bank with the name 'Monzo Bank Limited' in available banks")
	}

	// Flow requires user to click a link and authenticate with bank after running below code.
	// Simulating by authenticating manually and injecting env var of id returned into next function instead
	//req := nordigenAPI.CreateRequisition(accessBearerToken, bankId, "http://www.google.com")

	accounts := nordigenAPI.ListAccounts(accessBearerToken, os.Getenv("LINKED_REQ_ID"))
	fmt.Printf("%+v", accounts)

	transactions := nordigenAPI.ListTransactions(accessBearerToken, accounts.Accounts[0])

	fmt.Printf("%+v", transactions)

}
