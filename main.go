package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"gopkg.in/yaml.v2"
)

// TODO(#7): structure for other APIs
type Definition []struct {
	Meta struct {
		ID string `json:"id"`
	} `json:"meta"`

	Shortdef []string `json:"shortdef"`
}

func displayDef(definition []string, traverses int) {

	if traverses == (len(definition) - 1) {
		fmt.Printf("%d - %v", (traverses + 1), definition[traverses])
		return
	} else {
		fmt.Printf("%d - %v\n", (traverses + 1), definition[traverses])
		displayDef(definition, traverses+1)
	}
}

type config struct {
	Website    string `yaml:"website"`
	Link       string `yaml:"link`
	ApiKey     string `yaml:"apikey"`
	Dictionary string `yaml:"dictionary"`
}

func getConfig() (string, string, string, map[string][]string, error) {

	buf, err := ioutil.ReadFile("conf.yaml")
	if err != nil {
		return "", "", "", map[string][]string(nil), errors.New("conf.yaml - not in path")
	}

	conf := &config{}
	err = yaml.Unmarshal(buf, conf)
	if err != nil {
		return "", "", "", map[string][]string(nil), errors.New("conf.yaml - invalid configuration")
	}
	return conf.Website, conf.Link, conf.ApiKey, grabDict(conf.Dictionary), nil
}

func parseRequest(word string, website string, link string, apiKey string) (string, error) {

	// TODO(#6): add support for multiple dictionary apis
	switch website {
	case "dictionary.com":
		return fmt.Sprintf("%v%v%v", link, word, apiKey), nil
	}

	return "", errors.New("conf.Website invalid config")
}

func get(url string) Definition {

	//sponge
	fmt.Printf("%v\n", url)

	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)

	parsedReq := Definition{}
	json.Unmarshal(bodyBytes, &parsedReq)

	return parsedReq
	// displayDef(parsedReq[0].Shortdef, 0)
}

// TODO(#2): Implement more flags, is there a better way to parse flags?
func main() {

	if len(os.Args) < 2 {
		fmt.Printf("invalid number of arguments\n")
		return
	}

	website, link, apiKey, dictionary, err := getConfig()
	if err != nil {
		log.Fatalln(err)
	}
	
	definition, err := checkDict(os.Args[1], dictionary)
	if err == nil {
		displayDef(definition, 0)
		return
	}

	requestLink, err := parseRequest(os.Args[1], website, link, apiKey) 
	if err != nil {
		log.Fatalln(err)
	} else {
		// TODO(#8): implement a way to store already defined words, and check for them
		definition := get(requestLink)
		displayDef(definition[0].Shortdef, 0)
		//storeDef
	}
}
