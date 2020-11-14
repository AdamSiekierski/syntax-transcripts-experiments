package main

import (
	"fmt"
	"log"

	"github.com/AdamSiekierski/syntax-transcripts-experiments/utils"
)

type talkTime struct {
	wes   float64
	scott float64
}

func getTotalTalkTime(shows utils.Transcripts) talkTime {
	var tt talkTime

	for _, show := range shows {
		for _, sentence := range show {
			talkSeconds := sentence.End - sentence.Start

			talkMinutes := talkSeconds / 60

			if sentence.Initials == "WB" {
				tt.wes += talkMinutes
			}

			if sentence.Initials == "ST" {
				tt.scott += talkMinutes
			}
		}
	}

	return tt
}

func main() {
	transcripts, err := utils.LoadTranscripts()
	if err != nil {
		log.Fatal(fmt.Errorf("error loading transcript: %w", err))
	}

	totalTalkTime := getTotalTalkTime(transcripts)

	fmt.Printf("Total talk time: %.2f minutes, %.2f hours.\n", totalTalkTime.scott+totalTalkTime.wes, (totalTalkTime.scott+totalTalkTime.wes)/60)
	fmt.Printf("Total Wes talk time: %.2f minutes, %.2f hours.\n", totalTalkTime.wes, totalTalkTime.wes/60)
	fmt.Printf("Total Scott talk time: %.2f minutes, %.2f hours.\n", totalTalkTime.scott, totalTalkTime.scott/60)
}
