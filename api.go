package main

// TODO(#7): structure for other APIs
type definition []struct {
	Meta struct {
		ID string `json:"id"`
	} `json:"meta"`

	Shortdef []string `json:"shortdef"`
}
