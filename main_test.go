package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"testing"
)

func initConf() {

	API1, ok := os.LookupEnv("API1")
	if !ok {
		panic("$API1 - enviromental variable not set")
	}
	tmpBody := `
---
website:
 - api:
      website: "dictionaryapi.com"
      link: "https://www.dictionaryapi.com/api/v3/references/collegiate/json/"
      apikey: "` + API1 + `"
 - api:
      website: "api.dictionaryapi.dev"
      link: "https://api.dictionaryapi.dev/api/v2/entries/en/"
      apikey: NULL
dictionary: "dictionary.json"
`
	fmt.Println(tmpBody)
	err := ioutil.WriteFile("./conf.yaml", []byte(tmpBody), 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func getLogistics() (*config, string, map[string][]string) {
	initConf()
	apiConf, defPath := getConfig()
	dictFile := getDictConf(apiConf)
	dictionary := getDict(defPath + "/" + dictFile)

	return apiConf, dictFile, dictionary
}

func TestLocateDef(t *testing.T) {
	apiConf, dictFile, dictionary := getLogistics()

	var tests = []struct {
		a  string
		o1 bool
	}{
		{"undefinablestring", true},
		{"definable", false},
	}

	for _, tv := range tests {

		t.Run(tv.a, func(t *testing.T) {
			definition := locateDef(tv.a, apiConf, dictFile, dictionary)
			ok := verifyDef(definition)
			if ok != tv.o1 {
				t.Errorf("got: %s is a (%v) def - want (%v)", tv.a, ok, tv.o1)
			}
		})
	}
	os.Remove("conf.yaml")
}

func TestCallAPI(t *testing.T) {
	apiConf, _, _ := getLogistics()

	var tests = []struct {
		idx  int
		word string
		o1   int
		o2   int
	}{
		{0, "undefinableword", 0, 0},
		{0, "definable", 1, 10},
		{1, "undefinableword", 0, 0},
		{1, "definable", 1, 10},
	}

	for _, tv := range tests {

		website, link, apiKey := parseConfig(apiConf, tv.idx)
		t.Run(website, func(t *testing.T) {

			requestLink, err := parseRequest(tv.word, website, link, apiKey)
			if err != nil {
				t.Errorf("failed with a fatal error")
			} else {
				response := callAPI(website, requestLink)
				if (len(response) >= tv.o1) && (len(response) <= tv.o2) {
					return
				}

				t.Errorf("got %d - wanted %d/%d", len(response), tv.o1, tv.o2)
			}
		})
	}
	os.Remove("conf.yaml")
}
