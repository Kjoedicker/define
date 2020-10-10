package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

func toJSON(data map[string][]string) []byte {

	procData, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		log.Fatalf("error during json.Marshal\n")
	}
	return procData
}

func froJSON(unparsedJSON []byte) map[string][]string {

	var parsedJSON map[string][]string

	json.Unmarshal([]byte(unparsedJSON), &parsedJSON)

	return parsedJSON
}

func grabJSON(filename string) []byte {

	JSONFile, err := os.Open(filename)
	// TODO: add support for creating a dictionary file, if not provided already.
	if err != nil {
		log.Fatalf("grabJSON() - error in provided filepath\n")
	}
	byteValue, _ := ioutil.ReadAll(JSONFile)

	return byteValue
}

func storeJSON(filename string, data map[string][]string) {

	converted := toJSON(data)

	err := ioutil.WriteFile(filename, converted, 0644)
	if err != nil {
		panic(err)
	}
}
