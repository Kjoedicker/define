package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"gopkg.in/yaml.v2"
)

// TODO(#7): structure for other APIs
type definition []struct {
	Meta struct {
		ID string `json:"id"`
	} `json:"meta"`

	Shortdef []string `json:"shortdef"`
}

// TODO(#15): find other panic mode errors and handle them here
func recovery(errType string) {
	if r := recover(); r != nil {
		switch errType {
		case "invalid":
			fmt.Println("not in database")
		}
	}
}

type config struct {
	Website    string `yaml:"website"`
	Link       string `yaml:"link"`
	APIKey     string `yaml:"apikey"`
	Dictionary string `yaml:"dictionary"`
}

func getConfig() (string, string, string, string, error) {

	buf, err := ioutil.ReadFile("conf.yaml")
	if err != nil {
		return "", "", "", "", errors.New("conf.yaml - not in path")
	}

	conf := &config{}
	err = yaml.Unmarshal(buf, conf)
	if err != nil {
		return "", "", "", "", errors.New("conf.yaml - invalid configuration")
	}
	return conf.Website, conf.Link, conf.APIKey, conf.Dictionary, nil
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
		log.Fatalln(err)
	}

	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)

	parsedReq := definition{}
	json.Unmarshal(bodyBytes, &parsedReq)

	return parsedReq
}
