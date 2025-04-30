package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"golang-core/cli/utils"
)

var fullInfo bool
var meCmd *cobra.Command

func registerMeCommand() {
	meCmd = CreateCommand(
		"me",
		"Print the information about me.",
		"Print the full information about me.",
	)

	meCmd.Flags().BoolVarP(&fullInfo, "fullInfo", "f", false, "Print the full information about me.")
	meCmd.Run = func(cmd *cobra.Command, args []string) {
		printMe(fullInfo)
	}

	rootCmd.AddCommand(meCmd)
}

func printMe(fullInfo bool) {
	fmt.Print(utils.MeLogo)
	if !fullInfo {
		return
	}
	fmt.Print("\n✅ I am handsome.\n\n")
}
