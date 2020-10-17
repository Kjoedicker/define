package main

import (
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
	case "dictionaryapi.com":
		return fmt.Sprintf("%v%v%v", link, word, apiKey), nil
	case "api.dictionaryapi.dev":
		return fmt.Sprintf("%v%v", link, word), nil
	}

	return "", errors.New("conf.Website invalid config")
}

func get(url string) []byte {

	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Failed to resolve GET\n")
		log.Fatal(err)
	}

	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)

	return bodyBytes
}
