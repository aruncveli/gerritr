package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"golang.org/x/exp/slices"
)

var Branch string
var State string
var Message string

var supportedChangeStates = []string{"private", "remove-private", "wip",
	"ready"}

var rootCmd = &cobra.Command{
	Use:   "gerritr",
	Short: "Wrapping some Git for Gerrit",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		if State != "" {
			if !slices.Contains(supportedChangeStates, State) {
				fmt.Println("Ignoring unsupported state", State)
				State = ""
			}
		}
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&Branch, "branch", "b", "",
		"Target branch name")
	rootCmd.PersistentFlags().StringVarP(&State, "state", "s", "",
		"Change state: private, remove-private, wip or ready")
	rootCmd.PersistentFlags().StringVarP(&Message, "message", "m", "",
		"Commit message")
}
