package main

import (
	"errors"
	// "fmt"
)

/// TODO(#19): handle dictionary.json that is not made yet
func updateDict(dictionary map[string][]string, word string, definition []string) bool {
	if definition != nil {
		dictionary[word] = definition
		return false
	}

	return true
}

// TODO(#20): standardize where the dictionary will be kept
//   right now, things are kept in the same dir as the binary.
//   what are the standards people would want?
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
