package main

import (
	"fmt"
	"github.com/AdamSiekierski/syntax-experiments/utils"
	"strings"
)

func flattenTranscripts(shows utils.Transcripts, speakerInitials string) []string {
	var sentences []string

	for _, show := range shows {
		for _, sentence := range show {
			if speakerInitials == sentence.Initials || speakerInitials == "" {
				sentences = append(sentences, sentence.Text)
			}
		}
	}

	return sentences
}

func sentencesToString(sentences []string) string {
	return strings.Join(sentences, " ")
}

func main() {
	transcripts, err := utils.LoadTranscripts()

	if err != nil {
		fmt.Print(err)
	} else {
		allSentences := flattenTranscripts(transcripts, "")
		wesSentences := flattenTranscripts(transcripts, "WB")
		scottSentences := flattenTranscripts(transcripts, "ST")

		allSentencesString := sentencesToString(allSentences)
		wesSentencesString := sentencesToString(wesSentences)
		scottSentencesString := sentencesToString(scottSentences)

		allSicks := strings.Count(allSentencesString, "sick")
		wesSicks := strings.Count(wesSentencesString, "sick")
		scottSicks := strings.Count(scottSentencesString, "sick")

		fmt.Printf("How many times has been `sick` said? %d\n", allSicks)
		fmt.Printf("How many times has been `sick` said by Wes? %d\n", wesSicks)
		fmt.Printf("How many times has been `sick` said by Scott? %d\n", scottSicks)
	}
}
