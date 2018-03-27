package search

import (
	"os"
	"encoding/json"
)

const dataFile = "data/data.json"

type Feed struct {
	Name string `json:"name"`
	URI  string `json:"link"`
	Type string `json:"type"`
}

func RetrieveFeeds() ([]*Feed, error) {
	file, err := os.Open(dataFile)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var feeds []*Feed
	err = json.NewDecoder(file).Decode(&feeds)
	return feeds, err
}

