package storybuilder

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

func readJSON(jsonPath string) []byte {
	jsonFile, err := os.Open(jsonPath)

	if err != nil {
		panic(err)
	}

	jsonByteValue, err := ioutil.ReadAll(jsonFile)

	if err != nil {
		panic(err)
	}

	return jsonByteValue
}

// ArcOption describes an option to go to other arc
type ArcOption struct {
	Text    string `json:"text"`
	ArcLink string `json:"arc"`
}

// Arc is a single part of story
type Arc struct {
	Title      string      `json:"title"`
	Paragraphs []string    `json:"story"`
	Options    []ArcOption `json:"options"`
}

// ParseFromJSON Parses json to story type
func ParseFromJSON(jsonPath string) map[string]Arc {
	var story map[string]Arc

	jsonByteValue := readJSON(jsonPath)
	json.Unmarshal(jsonByteValue, &story)

	return story
}
