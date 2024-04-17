package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

var client = &http.Client{}

type hasteResponse struct {
	Key string `json:"key"`
}

func hastePost(data string) (*hasteResponse, error) {
	r, err := http.NewRequest("POST", fmt.Sprintf("%s/documents", *hasteURL), bytes.NewBuffer([]byte(data)))
	if err != nil {
		log.Println("Error creating request:", err.Error())
		return nil, err
	}

	// Set some headers
	r.Header.Set("Content-Type", "text/plain")
	r.Header.Set("Accept", "application/json")
	r.Header.Set("User-Agent", "EggLabs HasteCat v"+version)

	// Send the request
	resp, err := client.Do(r)
	if err != nil {
		log.Println("Error sending request:", err.Error())
		return nil, err
	}
	defer resp.Body.Close()

	// Check the response status code
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		log.Println("Error sending request: HTTP status code", resp.StatusCode)
		return nil, fmt.Errorf("HTTP status code %d", resp.StatusCode)
	}

	// Read and decode the response
	doc := &hasteResponse{}
	err = json.NewDecoder(resp.Body).Decode(doc)
	if err != nil {
		log.Println("Error decoding response:", err.Error())
		return nil, err
	}

	// Return the response
	return doc, nil

}
