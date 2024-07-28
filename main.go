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
	deckDirs, err := imageprocessing.FindAllEndDirsectories(Source)
	if err != nil {
		fmt.Println("error:", err)
	}

	decks, err := imageprocessing.LoadAllDecksDir(deckDirs)
	if err != nil {
		fmt.Println("error:", err)
	}

	imageprocessing.ExportDecks(decks, ResultDir)
}
