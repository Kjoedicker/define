package main

import (
	"encoding/json"
	"errors"
	// "fmt"
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

func updateDict(dictionary map[string][]string, word string, definition []string) {
	if dictionary[word] != nil {
		dictionary[word] = definition
	}
	
	// TODO(#18): Add a log for non critical errors that occur
	//	ex. the dictionary did not get written to because dictionary[word] is nil
}

func grabDict(dictionary string) map[string][]string {
	return froJSON(grabJSON(dictionary))
}

func checkDict(word string, dictionary map[string][]string) ([]string, error) {

	if word, ok := dictionary[word]; ok {
		return word, nil
	}

	return []string(nil), errors.New("checkDict() - undefined")
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
