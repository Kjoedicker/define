package main

import (
	"encoding/json"
)

// TODO(#7): structure for other APIs
type merriamAPI []struct {
	Meta struct {
		ID string `json:"id"`
	} `json:"meta"`

	Shortdef []string `json:"shortdef"`
}

func (API merriamAPI) getDef() []string {
	defer recovery("invalid")
	return API[0].Shortdef
}

func (API merriamAPI) marshallAPI(APIType string, bodyBytes []byte) merriamAPI {
	parsedReq := merriamAPI{}
	json.Unmarshal(bodyBytes, &parsedReq)
	return parsedReq
}

type googleAPI []struct {
	Word      string `json:"word"`
	Phonetics []struct {
		Text  string `json:"text"`
		Audio string `json:"audio"`
	} `json:"phonetics"`
	Meanings []struct {
		PartOfSpeech string `json:"partOfSpeech"`
		Definitions  []struct {
			Definition string   `json:"definition"`
			Example    string   `json:"example"`
			Synonyms   []string `json:"synonyms"`
		} `json:"definitions"`
	} `json:"meanings"`
}

// TODO(#27): does not properly handle nested definitions for unGoogleAPI
func (API googleAPI) getDef() []string {

	if len(API) > 0 {
		definitions := make([]string, len(API[0].Meanings))
		for idx := range API[0].Meanings {
			definitions[idx] = API[0].Meanings[idx].Definitions[0].Definition
		}
		return definitions
	}
	return []string(nil)
}

func (API googleAPI) marshallAPI(APIType string, bodyBytes []byte) googleAPI {
	parsedReq := googleAPI{}
	json.Unmarshal(bodyBytes, &parsedReq)
	return parsedReq
}

func callAPI(website string, requestLink string) []string {

	switch website {
	case "dictionaryapi.com":
		parsedReq := merriamAPI{}
		got := parsedReq.marshallAPI(website, get(requestLink))
		return got.getDef()

	case "api.dictionaryapi.dev":
		parsedReq := googleAPI{}
		got := parsedReq.marshallAPI(website, get(requestLink))
		return got.getDef()
	}

	return []string(nil)
}
