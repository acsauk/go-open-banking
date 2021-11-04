package main

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/acsauk/go-open-banking/internal/nordigenAPI"
	"github.com/joho/godotenv"
)

func main() {
	// load .env file
	err := godotenv.Load()

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	accessBearerToken := nordigenAPI.GetBearerAccessToken()

	// Get available banks
	client := http.Client{}
	req, err := http.NewRequest("GET", "https://ob.nordigen.com/api/v2/institutions/?country=gb", nil)
	if err != nil {
		//Handle Error
	}

	log.Println(accessBearerToken)

	req.Header = http.Header{
		"Host":          []string{"www.host.com"},
		"Content-Type":  []string{"application/json"},
		"Authorization": []string{accessBearerToken},
	}

	resp, err := client.Do(req)
	if err != nil {
		//Handle Error
	}

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	log.Println(string(body))
}
