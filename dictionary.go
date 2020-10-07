package main

import (
	"errors"
	// "fmt"
)

func updateDict(dictionary map[string][]string, word string, definition []string) bool {
	if definition != nil {
		dictionary[word] = definition
		return false
	} else {
		return true
	}
	
	// TODO(#18): Add a log for non critical errors that occur
	//	 ex. the dictionary did not get written to because dictionary[word] is nil
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
