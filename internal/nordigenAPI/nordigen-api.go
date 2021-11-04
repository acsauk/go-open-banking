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
