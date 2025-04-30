package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"golang-core/cli/utils"
)

var timeFormat string
var timeNowCmd *cobra.Command

func registerTimeNowCommand() {
	weatherCmd = CreateCommand(
		"timenow",
		"Print timenow.",
		"Print timenow.",
	)

	weatherCmd.Flags().StringVarP(
		&timeFormat, "format", "f", "",
		"Format output for timenow.\nAvailable formats: "+utils.AvailableTimeFormats())
	weatherCmd.RunE = func(cmd *cobra.Command, args []string) error {
		timeNow, err := utils.TimeNow(timeFormat)
		if err != nil {
			return err
		}
		fmt.Println(timeNow)
		return nil
	}

	rootCmd.AddCommand(weatherCmd)
}
