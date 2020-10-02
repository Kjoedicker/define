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

// TODO(#11): flags to specify display type. json, text,?
func displayDef(definition []string, traverses int) {

	if traverses == (len(definition) - 1) {
		fmt.Printf("%d - %v\n", (traverses + 1), definition[traverses])
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

func getConfig() (string, string, string, string, error) {

	buf, err := ioutil.ReadFile("conf.yaml")
	if err != nil {
		return "", "", "", "", errors.New("conf.yaml - not in path\n")
	}

	conf := &config{}
	err = yaml.Unmarshal(buf, conf)
	if err != nil {
		return "", "", "", "", errors.New("conf.yaml - invalid configuration\n")
	}
	return conf.Website, conf.Link, conf.ApiKey, conf.Dictionary, nil
}

func parseRequest(word string, website string, link string, apiKey string) (string, error) {

	// TODO(#6): add support for multiple dictionary apis
	switch website {
	case "dictionary.com":
		return fmt.Sprintf("%v%v%v", link, word, apiKey), nil
	}

	return "", errors.New("conf.Website invalid config\n")
}

// TODO(#12): add handling for words not found during api request
func get(url string) Definition {

	//sponge
	fmt.Printf("referencing api..\n")

	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)

	parsedReq := Definition{}
	json.Unmarshal(bodyBytes, &parsedReq)

	return parsedReq
}

func procWord(word string) {

	// TODO: add a flag for displaying the word before the definition
	fmt.Printf("\n%v:\n", word)

	website, link, apiKey, dictFile, err := getConfig()
	if err != nil {
		log.Fatalln(err)
	}

	dictionary := grabDict(dictFile)
	definition, err := checkDict(word, dictionary)
	if err == nil {
		displayDef(definition, 0)
		return
	}

	requestLink, err := parseRequest(word, website, link, apiKey) 
	if err != nil {
		log.Fatalln(err)
	} else {
		definition := get(requestLink)
		displayDef(definition[0].Shortdef, 0)
		updateDict(dictionary, word, definition[0].Shortdef)
		storeJson(dictFile, dictionary)
	}
}

// TODO(#2): Implement more flags, is there a better way to parse flags?
func main() {

	if len(os.Args) < 2 {
		fmt.Printf("invalid number of arguments\n")
		return
	} else {
		for index := 1; index < len(os.Args); index++ {
			procWord(os.Args[index])
		}
	}
}
