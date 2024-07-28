package clicommands

import (
	"fmt"
	"os"
	imageprocessing "tts-deck-gen/image-processing"

	"github.com/spf13/cobra"
)

type generateAutoLocateParams struct {
	searchDir string
	exportDir string
}

var galParams = &generateAutoLocateParams{}

func GenerateAutoLocateCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "auto-locate",
		Short:   "Generate decks using provided search directory",
		PreRunE: generateAutoLocateValidateParams,
		Run:     generateAutoLocateCommand,
	}

	cmd.Flags().StringVar(
		&galParams.searchDir,
		"searchDir",
		"",
		"Path to directory to be searched",
	)

	cmd.Flags().StringVar(
		&galParams.exportDir,
		"exportDir",
		"",
		"Path to directory where decks will be expoted",
	)

	return cmd
}

func generateAutoLocateCommand(cmd *cobra.Command, _ []string) {
	deckDirs, err := imageprocessing.FindAllEndDirsectories(galParams.searchDir)
	if err != nil {
		fmt.Printf("Error locating directories %v\n", err)
		return
	}

	decks, err := imageprocessing.LoadAllDecksDir(deckDirs)
	if err != nil {
		fmt.Printf("Error loading decks %v\n", err)
		return
	}

	err = imageprocessing.ExportDecks(decks, galParams.exportDir)
	if err != nil {
		fmt.Printf("Error exporting decks %v\n", err)
		return
	}
}

func generateAutoLocateValidateParams(_ *cobra.Command, _ []string) error {
	fileInfo, err := os.Stat(galParams.searchDir)
	if err != nil {
		if os.IsNotExist(err) {
			return fmt.Errorf("searchDir doesn't exist: %v", err)
		}
		return fmt.Errorf("unknown error: %v", err)
	}

	if !fileInfo.IsDir() {
		return fmt.Errorf("searchDir is not a directory")
	}

	fileInfo, err = os.Stat(galParams.exportDir)
	if err != nil {
		if os.IsNotExist(err) {
			return fmt.Errorf("exportDir doesn't exist: %v", err)
		}
		return fmt.Errorf("unknown error: %v", err)
	}

	if !fileInfo.IsDir() {
		return fmt.Errorf("exportDir is not a directory")
	}

	return nil
}
