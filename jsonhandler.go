package main

import (
	"encoding/json"
	"errors"
	// "fmt"
	"io/ioutil"
	"log"
	"os"
)

func toJson(data map[string][]string) []byte {

	procData, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		log.Fatalf("error during json.Marshal\n")
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
		log.Fatalf("grabJson() - error in provided filepath\n")
	}
	byteValue, _ := ioutil.ReadAll(jsonFile)

	return byteValue
}

func storeJson(filename string, data map[string][]string) {

	converted := toJson(data)

	err := ioutil.WriteFile(filename, converted, 0644)
	if err != nil {
		panic(err)
	}
}

func updateDict(dictionary map[string][]string, word string, definition []string) {
	dictionary[word] = definition
}

func grabDict(dictionary string) map[string][]string {
	return froJson(grabJson(dictionary))
}

func checkDict(word string, dictionary map[string][]string) ([]string, error) {

	if word, ok := dictionary[word]; ok {
		return word, nil
	}

	return []string(nil), errors.New("checkDict() - undefined\n")
}

// func main() {

// 	// test := froJson(grabJson("mesh.json"))
// 	if l, err := checkDict("test1"); err != nil {
// 		fmt.Printf("here")
// 	} else {
// 		fmt.Printf("%v", l)
// 	}

// 	// fmt.Printf("%v", test[os.Args[1]])
// 	// fmt.Printf("%v", checkDict("test"))
// }
