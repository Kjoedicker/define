package main

import (
	"errors"
	// "fmt"
)

func updateDict(dictionary map[string][]string, word string, definition []string) bool {
	if definition != nil {
		dictionary[word] = definition
		return false
	}
	return true
}

func getDict(dictionary string) map[string][]string {
	return froJSON(grabJSON(dictionary))
}

func checkDict(word string, dictionary map[string][]string) ([]string, error) {

	if word, ok := dictionary[word]; ok {
		return word, nil
	}

	return []string(nil), errors.New("checkDict() - undefined")
}
