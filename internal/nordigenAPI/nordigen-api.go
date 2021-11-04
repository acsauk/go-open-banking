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

func GetAvailableBanks(accessBearerToken string) []Bank {
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
		//Handle Error
	}

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

	//pretty := &bytes.Buffer{}
	//if err := json.Indent(pretty, banksJson, "", "  "); err != nil {
	//	panic(err)
	//}
	//
	//log.Println(pretty.String())
	return banks
}
