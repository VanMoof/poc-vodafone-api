package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func main() {
	// staging base url
	baseURL := "https://api.m2m.vodafone.com/m2m/v2/oauth2/access-token"
	u := url.Values{
		"grant_type": []string{""},
		"scope":      []string{""},
		"username":   []string{""},
		"password":   []string{""},
		"customerId": []string{""},
	}

	urlEncoded := u.Encode()

	client := &http.Client{}

	responseBody := bytes.NewBuffer([]byte(urlEncoded))

	req, err := http.NewRequest("POST", baseURL, responseBody)
	if err != nil {
		fmt.Println("New request went wrong")
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Authorization", "")
	fmt.Println(req.Body)
	resp, err := client.Do(req)
	if err != nil {
		fmt.Print("Something went wrong")
	}

	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))
}
