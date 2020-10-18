package main

import (
	"errors"
	// "fmt"
)

func updateDict(dictionary map[string][]string, word string, definition []string) bool {
	if len(definition) > 0 {
		dictionary[word] = definition
		return false
	}

	return true
}

func getDict(dictionary string) map[string][]string {

	tmpDict := froJSON(grabJSON(dictionary))
	if tmpDict != nil {
		return tmpDict
	}

	return make(map[string][]string)
}

func checkDict(word string, dictionary map[string][]string) ([]string, error) {

	if word, ok := dictionary[word]; ok {
		return word, nil
	}

	return []string(nil), errors.New("checkDict() - undefined")
}
