package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// TODO(#15): find other panic mode errors and handle them here
func recovery(errType string) {
	if r := recover(); r != nil {
		switch errType {
		case "invalid":
			fmt.Println("not in database")
		}
	}
}

func parseRequest(word string, website string, link string, apiKey string) (string, error) {

	// TODO(#6): add support for multiple dictionary apis
	switch website {
	case "dictionary.com":
		return fmt.Sprintf("%v%v%v", link, word, apiKey), nil
	}

	return "", errors.New("conf.Website invalid config")
}

func get(url string) definition {

	//sponge
	// fmt.Printf("referencing api..\n%v", url)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Failed to resolve GET\n")
		log.Fatal(err)
	}

	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)

	parsedReq := definition{}
	json.Unmarshal(bodyBytes, &parsedReq)

	return parsedReq
}
