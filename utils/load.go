package utils

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type timestamp struct {
	HH string `json:"hh"`
	MM string `json:"mm"`
	SS string `json:"ss"`
}

// Sentence - Single sentence from JSON
type Sentence struct {
	ID        int       `json:"id"`
	Text      string    `json:"text"`
	Speaker   string    `json:"speaker"`
	Initials  string    `json:"initials"`
	Start     float64   `json:"start"`
	End       float64   `json:"end"`
	Timestamp timestamp `json:"timestamp"`
}

// Show - A set of sentences
type Show []Sentence

// Transcripts - A set of all shows
type Transcripts []Show

func loadTranscriptsFile() ([]byte, error) {
	file, err := os.Open("./transcripts.json")

	if err != nil {
		return nil, err
	}

	defer file.Close()

	byteArray, err := ioutil.ReadAll(file)

	if err != nil {
		return nil, err
	}

	return byteArray, nil
}

// LoadTranscripts - Load all transcripts from JSON file
func LoadTranscripts() (Transcripts, error) {
	file, err := loadTranscriptsFile()

	if err != nil {
		return nil, err
	}

	var transcripts Transcripts

	err = json.Unmarshal(file, &transcripts)

	if err != nil {
		return nil, err
	}

	return transcripts, nil
}
