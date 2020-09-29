package main

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
	"log"
	"os"
    "net/http"
)

type Definition []struct {

	Meta struct {
		ID string `json:"id"`	
	} `json:"meta"`

	Shortdef []string `json:"shortdef"`
}

func display(Shortdef []string, traverses int) {

	if (traverses == (len(Shortdef)-1)) {
		fmt.Printf("%d - %v", (traverses+1), Shortdef[traverses])
		return 
	} else {
		fmt.Printf("%d - %v\n", (traverses+1), Shortdef[traverses])
		display(Shortdef, traverses+1)
	}
}

// TODO(#3): Implement an getKey() func, that will look for an api key in conf.yaml
// getKey() string {}

// TODO(#4): grab data from a conf file to produce a request
func concatRequest(word string) string {
	
	// TODO(#5): Replace dev key with getKey() func 
	return fmt.Sprintf("%v%v%v","https://www.dictionaryapi.com/api/v3/references/collegiate/json/", word, "?key=2725bb6b-51ac-41c9-a400-3b863c04cca5")
}

// TODO(#1): implement error handling on api
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
	// display(parsedReq[0].Shortdef, 0)
}

// TODO(#2): Implement more flags, is there a better way to parse flags?
func main() {

	if (len(os.Args) == 2) {
		definition := get(concatRequest(os.Args[1]))
		display(definition[0].Shortdef, 0)
	} else {
		fmt.Printf("invalid arguments")
	}
}
