package main

import (
	"fmt"
	imageprocessing "tts-deck-gen/image-processing"
)

const (
	Source    = "E:\\Project\\Y\\Final\\A-Decks"
	ResultDir = "E:\\Project\\Y\\Final\\Generated"
)

func main() {
	decks, err := imageprocessing.LoadAllDecks(Source)
	if err != nil {
		fmt.Println("error:", err)
	}

	decks.ExportDecks(ResultDir)
}
