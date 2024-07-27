package main

import (
	"fmt"
	imageprocessing "tts-deck-gen/image-processing"
)

func main() {
	decks, err := imageprocessing.LoadAllDecks("E:\\Project\\Y\\Final\\Decks")
	if err != nil {
		fmt.Println("FUCK MY LIFE", err)
	}

	for _, deck := range decks.Decks {
		fmt.Println(deck.DirPath)
	}
}
