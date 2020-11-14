package main

import (
	"fmt"
	"github.com/AdamSiekierski/syntax-transcripts-experiments/utils"
	"log"
	"os"
	"regexp"
	"sort"
	"strings"
)

func splitSentences(shows utils.Transcripts, speakerInitials string) []string {
	var result []string

	// Remove commas, dots, etc.
	m := regexp.MustCompile("[.,:]")

	for _, show := range shows {
		for _, sentence := range show {
			if sentence.Initials == speakerInitials || speakerInitials == "" {
				sentence.Text = m.ReplaceAllString(sentence.Text, "")

				result = append(result, strings.Split(strings.ToLower(sentence.Text), " ")...)
			}
		}
	}

	return result
}

type word struct {
	Place       int
	Word        string
	Occurencies int
}

func getOccurencies(words []string) []word {
	occurencies := make(map[string]int)

	for _, word := range words {
		_, ok := occurencies[word]

		if ok {
			occurencies[word]++
		} else {
			occurencies[word] = 1
		}
	}

	var occurenciesSlice []word

	for key, value := range occurencies {
		occurenciesSlice = append(occurenciesSlice, word{Word: key, Occurencies: value, Place: 0})
	}

	sort.Slice(occurenciesSlice, func(i, j int) bool {
		return occurenciesSlice[i].Occurencies > occurenciesSlice[j].Occurencies
	})

	for i, word := range occurenciesSlice {
		word.Place = i + 1
		occurenciesSlice[i] = word
	}

	return occurenciesSlice
}

func main() {
	transcripts, err := utils.LoadTranscripts()

	if err != nil {
		log.Fatal(fmt.Errorf("error loading transcript: %w", err))
	}

	allWords := splitSentences(transcripts, "")
	allOccurencies := getOccurencies(allWords)

	scottWords := splitSentences(transcripts, "ST")
	scottOccurencies := getOccurencies(scottWords)

	wesWords := splitSentences(transcripts, "WB")
	wesOccurencies := getOccurencies(wesWords)

	type jsonFile struct {
		All   []word `json:"all"`
		Scott []word `json:"scott"`
		Wes   []word `json:"wes"`
	}

	f, err := os.Create("./experiments/occurencies/occurencies.json")
	if err != nil {
		log.Fatal(fmt.Errorf("error creating json file: %w", err))
	}

	defer f.Close()

	jsonString := utils.Format(jsonFile{All: allOccurencies[0:100], Scott: scottOccurencies[0:100], Wes: wesOccurencies[0:100]})

	f.WriteString(jsonString)

	fmt.Printf("How many words have been used? %d\n", len(allOccurencies))
	fmt.Printf("How many words has Scott used? %d\n", len(scottOccurencies))
	fmt.Printf("How many words has Wes used?   %d\n", len(wesOccurencies))

	fmt.Println("Check out experiments/occurencies/occurencies.json for 100 most commonly used words!")
}
