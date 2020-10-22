package main

func verifyDef(definition []string) bool {

	if len(definition) > 0 {
		return false
	}

	return true
}

func updateDict(dictionary map[string][]string, word string, definition []string) {
	dictionary[word] = definition
}

func getDict(dictionary string) map[string][]string {

	tmpDict := froJSON(grabJSON(dictionary))
	if tmpDict != nil {
		return tmpDict
	}

	return make(map[string][]string)
}

func checkDict(word string, dictionary map[string][]string) ([]string, bool) {

	if word, ok := dictionary[word]; ok {
		return word, true
	}

	return []string(nil), false
}
