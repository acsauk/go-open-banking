package nordigenAPI

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type Bank struct {
	Name string `json:"name"`
	Id string `json:"id"`
}

type Requisition struct {
	Id string `json:"id"`
	Redirect string `json:"redirect"`
	Status string `json:"status"`
	Agreements string `json:"agreements"`
	Link string `json:"link"`
}

type Accounts struct {
	Accounts []string `json:"accounts"`
}

type TransactionAmount struct {
	Amount string `json:"amount"`
	Currency string `json:"currency"`
}

type BookedTransaction struct {
	Date string `json:"bookingDate"`
	CreditorName string `json:"creditorName"`
	TransactionAmount TransactionAmount `json:"transactionAmount"`
}

type Transactions struct {
	BookedTransactions []BookedTransaction `json:"booked"`
}

type wrapper struct {
	Transactions Transactions `json:"transactions"`
}

func GetBearerAccessToken() string {
	// Get access token
	body, err := json.Marshal(map[string]string{
		"secret_id":  os.Getenv("SECRET_ID"),
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

	return fmt.Sprintf("Bearer %s", t.AccessToken)
}

func GetAvailableBanks() []Bank {
	accessBearerToken := GetBearerAccessToken()

	// Get available banks
	client := http.Client{}
	req, err := http.NewRequest("GET", "https://ob.nordigen.com/api/v2/institutions/?country=gb", nil)
	if err != nil {
		//Handle Error
	}

	req.Header = http.Header{
		"Content-Type":  []string{"application/json"},
		"Authorization": []string{accessBearerToken},
	}

	resp, err := client.Do(req)

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	banksJson, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	var banks []Bank

	err = json.Unmarshal(banksJson, &banks)

	if err != nil {
		log.Fatal(err)
	}

	return banks
}

func CreateRequisition(instituteId, redirectURI string) Requisition {
	accessBearerToken := GetBearerAccessToken()

	body, err := json.Marshal(map[string]string{
		"redirect":  redirectURI,
		"institution_id": instituteId,
	})

	client := http.Client{}
	req, err := http.NewRequest("POST", "https://ob.nordigen.com/api/v2/requisitions/",  bytes.NewBuffer(body))
	if err != nil {
		//Handle Error
	}

	req.Header = http.Header{
		"Content-Type":  []string{"application/json"},
		"Authorization": []string{accessBearerToken},
	}

	resp, err := client.Do(req)

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	reqJSON, err := ioutil.ReadAll(resp.Body)

	var requisition Requisition

	err = json.Unmarshal(reqJSON, &requisition)

	if err != nil {
		log.Fatal(err)
	}

	return requisition
}

func ListAccounts(reqId string) Accounts {
	accessBearerToken := GetBearerAccessToken()

	client := http.Client{}
	uri := fmt.Sprintf("https://ob.nordigen.com/api/v2/requisitions/%s/", reqId)
	req, err := http.NewRequest("GET", uri, nil)

	if err != nil {
		//Handle Error
	}

	req.Header = http.Header{
		"Content-Type":  []string{"application/json"},
		"Authorization": []string{accessBearerToken},
	}

	resp, err := client.Do(req)

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	reqJSON, err := ioutil.ReadAll(resp.Body)

	var accounts Accounts

	err = json.Unmarshal(reqJSON, &accounts)

	if err != nil {
		log.Fatal(err)
	}

	return accounts
}

func ListTransactions(accountId string) []BookedTransaction {
	accessBearerToken := GetBearerAccessToken()

	client := http.Client{}
	uri := fmt.Sprintf("https://ob.nordigen.com/api/v2/accounts/%s/transactions/", accountId)
	req, err := http.NewRequest("GET", uri, nil)

	if err != nil {
		//Handle Error
	}

	req.Header = http.Header{
		"Content-Type":  []string{"application/json"},
		"Authorization": []string{accessBearerToken},
	}

	resp, err := client.Do(req)

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	transactionsJSON, err := ioutil.ReadAll(resp.Body)

	var wrapper wrapper

	err = json.Unmarshal(transactionsJSON, &wrapper)

	if err != nil {
		log.Fatal(err)
	}

	return wrapper.Transactions.BookedTransactions
}
