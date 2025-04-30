package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "mrroot501",
	Short: "mrroot501 is a cli tool.",
	Long:  "mrroot501 is a cli tool.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	rootCmd.CompletionOptions.DisableDefaultCmd = true
	registerAllCommands()
}

func registerAllCommands() {
	registerMeCommand()
	registerTimeNowCommand()
	registerWeatherCommand()
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: '%s'\n", err)
		os.Exit(1)
	}
}

func CreateCommand(use, short, long string) *cobra.Command {
	return &cobra.Command{
		Use:   use,
		Short: short,
		Long:  long,
	}
}
