package main

import (
	"fmt"
	"os"
	clicommands "tts-deck-gen/cli-commands"

	"github.com/spf13/cobra"
)

const (
	Source    = "E:\\Project\\Y\\Final\\A-Decks"
	ResultDir = "E:\\Project\\Y\\Final\\Generated"
)

func main() {

	var rootCmd = &cobra.Command{
		Use:   "tts-deck-gen",
		Short: "A brief description of your application",
		Long:  `A longer description of your application.`,
	}

	rootCmd.AddCommand(clicommands.GenerateAutoLocateCommand())
	rootCmd.AddCommand(clicommands.GenerateWithConfigCommand())

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
