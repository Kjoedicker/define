package main

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
	"log"
	"os"
	"net/http"
	"errors"
	"gopkg.in/yaml.v2"
)

// TODO: structure for other APIs
type Definition []struct {

	Meta struct {
		ID string `json:"id"`	
	} `json:"meta"`

	Shortdef []string `json:"shortdef"`
}

func displayDef(Shortdef []string, traverses int) {

	if (traverses == (len(Shortdef)-1)) {
		fmt.Printf("%d - %v", (traverses+1), Shortdef[traverses])
		return 
	} else {
		fmt.Printf("%d - %v\n", (traverses+1), Shortdef[traverses])
		displayDef(Shortdef, traverses+1)
	}
}

type config struct {
	Website string 	`yaml:"website"`
	Link 	string 	`yaml:"link`
	ApiKey 	string 	`yaml:"apikey"`
}

func getConfig() (string, string, string, error) {

	buf, err := ioutil.ReadFile("conf.yaml")
	if err != nil {
		return "", "", "", errors.New("conf.yaml - not in path")
	}

	conf := &config{}
	err = yaml.Unmarshal(buf, conf)
	if err != nil {
		return "", "", "", errors.New("conf.yaml - invalid configuration")
	}

	return conf.Website, conf.Link, conf.ApiKey, nil
}

func parseRequest(word string) (string, error) {

	website, link, apiKey, err := getConfig()
	if err != nil {
		log.Fatalln(err)
	} else {
		// TODO(#6): add support for multiple dictionary apis
		switch (website) {
		case "dictionary.com":
			return fmt.Sprintf("%v%v%v", link, word, apiKey), nil
		}
	}
	return "", errors.New("conf.yaml - invalid arguments supplied")
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

	if (len(os.Args) == 2) {
		link, err := parseRequest(os.Args[1])
		if err != nil {
			log.Fatalln(err)
		} else {
			definition := get(link)
			displayDef(definition[0].Shortdef, 0)
		}
	} else {
		fmt.Printf("invalid arguments")
	}
}
