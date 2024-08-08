package clicommands

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	imageprocessing "tts-deck-gen/image-processing"

	"github.com/spf13/cobra"
)

type generateWithConfigParams struct {
	configPath string
}

var gwcParams = &generateWithConfigParams{}

func GenerateWithConfigCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "with-config",
		Short:   "Generate decks using provided .json configuration",
		PreRunE: generateWithConfigValidateParams,
		Run:     generateWithConfigCommand,
	}

	cmd.Flags().StringVar(
		&gwcParams.configPath,
		"config-path",
		"",
		"JSON Config file to be parsed",
	)

	return cmd
}

func generateWithConfigCommand(cmd *cobra.Command, _ []string) {
	config, err := imageprocessing.ParseFromJson(gwcParams.configPath)
	if err != nil {
		fmt.Printf("Error parsing json config %v\n", err)
		return
	}

	decks, err := imageprocessing.LoadAllDecksConfig(config)
	if err != nil {
		fmt.Printf("Error loading decks %v\n", err)
		return
	}

	err = imageprocessing.ExportDecks(decks, config.ExportPath)
	if err != nil {
		fmt.Printf("Error exporting decks %v\n", err)
		return
	}
}

func generateWithConfigValidateParams(_ *cobra.Command, _ []string) error {
	// Check if the file exists
	fileInfo, err := os.Stat(gwcParams.configPath)
	if err != nil {
		if os.IsNotExist(err) {
			return fmt.Errorf("file doesn't exist: %v", err)
		}
		return fmt.Errorf("unknown error: %v", err)
	}

	// Check if it is a file (not a directory)
	if fileInfo.IsDir() {
		return fmt.Errorf("file is a directory")
	}

	// Optionally, check if the file has a .json extension
	if filepath.Ext(gwcParams.configPath) != ".json" {
		return fmt.Errorf("file is not json")
	}

	// Read the file content
	content, err := os.ReadFile(gwcParams.configPath)
	if err != nil {
		return fmt.Errorf("file not readable %v", err)
	}

	// Try to parse the content as JSON
	var jsonData interface{}
	if err := json.Unmarshal(content, &jsonData); err != nil {
		return fmt.Errorf("invalid json")
	}

	return nil
}
