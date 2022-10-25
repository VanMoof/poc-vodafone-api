package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	// staging base url
	baseURL := "https://api.m2m.vodafone.com/m2m/v2/oauth2/access-token"
	params, _ :=
		json.Marshal(map[string]string{
			"grant_type": "password",
			"scope":      "M2M_NETWORK_ALL",
			"username":   "",
			"password":   "",
			"customerId": "",
		})

	responseBody := bytes.NewBuffer(params)
	req, err := http.NewRequest("POST", baseURL, responseBody)
	if err != nil {
		log.Fatalf("An Error Occured %v", err)
	}
	defer req.Body.Close()
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "")

	body, err := ioutil.ReadAll(req.Body)

	if err != nil {
		log.Fatalln(err)
	}

	fmt.Print(json.Unmarshal(body, &req))
}
