package main

import (
	"testing"
)

func getLogistics() (*config, string, string, map[string][]string) {
	apiConf, defPath := getConfig()
	dictFile := getDictConf(apiConf)
	dictionary := getDict(defPath + "/" + dictFile)

	return apiConf, defPath, dictFile, dictionary
}

func TestLocateDef(t *testing.T) {
	apiConf, defPath, dictFile, dictionary := getLogistics()

	var tests = []struct {
		a  string
		o1 int
	}{
		{"undefinablestring", 1},
		{"definable", 0},
	}

	for _, tv := range tests {

		t.Run(tv.a, func(t *testing.T) {
			_, ok := locateDef(tv.a, apiConf, defPath, dictFile, dictionary)
			if ok == tv.o1 {
				t.Errorf("got %d - want %d", ok, tv.o1)
			}
		})
	}
}

func TestCallAPI(t *testing.T) {
	apiConf, _, _, _ := getLogistics()

	var tests = []struct {
		idx  int
		word string
		o1   int
		o2   int
	}{
		{0, "undefinableword", 0, 0},
		{0, "definable", 1, 10},
		{1, "undefinableword", 1, 0},
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
}
