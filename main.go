package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	// load .env file
	err := godotenv.Load()

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// Get access token
	body, err := json.Marshal(map[string]string{
		"secret_id": os.Getenv("SECRET_ID"),
		"secret_key": os.Getenv("SECRET_KEY"),
	})

	if err != nil {
		log.Fatal(err)
	}

	resp, err := http.Post("https://ob.nordigen.com/api/v2/token/new/", "application/json", bytes.NewBuffer(body))

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	body, err = ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	type tokenResponse struct {
		AccessToken string `json:"access"`
	}

	t := tokenResponse{}
	json.Unmarshal(body, &t)

	token := fmt.Sprintf("Bearer %s", t.AccessToken)

	// Get available banks
	client := http.Client{}
	req , err := http.NewRequest("GET", "https://ob.nordigen.com/api/v2/institutions/?country=gb", nil)
	if err != nil {
		//Handle Error
	}

	log.Println(token)

	req.Header = http.Header{
		"Host": []string{"www.host.com"},
		"Content-Type": []string{"application/json"},
		"Authorization": []string{token},
	}

	resp, err = client.Do(req)
	if err != nil {
		//Handle Error
	}

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	body, err = ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	log.Println(string(body))

}