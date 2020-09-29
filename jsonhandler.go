package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func toJson(data map[string][]string) []byte {

	procData, err := json.Marshal(data)
	if err != nil {
		log.Fatalf("error during json.Marshal")
	}
	return procData
}

func froJson(unparsedJson []byte) map[string][]string {

	var parsedJson map[string][]string

	json.Unmarshal([]byte(unparsedJson), &parsedJson)

	return parsedJson
}

func grabJson(filename string) []byte {

	jsonFile, err := os.Open(filename)
	if err != nil {
		log.Fatalf("error in provided filepath")
	}
	byteValue, _ := ioutil.ReadAll(jsonFile)

	return byteValue
}

func checkDict(word string) ([]string, error) {
	dictionary := froJson(grabJson("mesh.json"))

	fmt.Printf("%v", dictionary[word])
	if word, ok := dictionary[word]; ok {
		return word, nil
	}

	return []string(nil), errors.New("undefined")
}

func main() {

	// test := froJson(grabJson("mesh.json"))
	if l, err := checkDict("test1"); err != nil {
		fmt.Printf("here")
	} else {
		fmt.Printf("%v", l)
	}

	// fmt.Printf("%v", test[os.Args[1]])
	// fmt.Printf("%v", checkDict("test"))
}
