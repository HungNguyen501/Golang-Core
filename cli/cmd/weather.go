package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"golang-core/cli/utils"
)

var weatherCmd *cobra.Command

func registerWeatherCommand() {
	weatherCmd = CreateCommand(
		"weather",
		"Print weather information.",
		"Print weather information.",
	)

	weatherCmd.RunE = func(cmd *cobra.Command, args []string) error {
		output, err := utils.Weather()
		if err != nil {
			return err
		}
		fmt.Println(output)
		return nil
	}

	rootCmd.AddCommand(weatherCmd)
}
