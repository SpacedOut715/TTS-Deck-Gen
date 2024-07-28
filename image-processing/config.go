package imageprocessing

import (
	"encoding/json"
	"fmt"
	"os"
)

type DeckLocation struct {
	DeckPath     string `json:deckPath`
	DeckFileName string `json:deckFileName`
}

type DecksConfig struct {
	Decks      []DeckLocation `json:decks`
	ExportPath string         `json:exportPath`
}

func ParseFromJson(configPath string) (*DecksConfig, error) {
	content, err := os.ReadFile(configPath)
	if err != nil {
		return nil, err
	}

	var jsonData DecksConfig
	if err := json.Unmarshal(content, &jsonData); err != nil {
		return nil, err
	}

	fmt.Printf("Config valid. Generating decks...\n")

	return &jsonData, nil
}
