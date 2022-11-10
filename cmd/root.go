package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"golang.org/x/exp/slices"
)

// Target branch name
var Branch string

/*
Change state

Ignored if the value is not one of [private, remove-private], [wip or ready]

[private, remove-private]: https://gerrit-documentation.storage.googleapis.com/Documentation/3.6.2/user-upload.html#private
[wip or ready]: https://gerrit-documentation.storage.googleapis.com/Documentation/3.6.2/user-upload.html#wip
*/
var State string
var supportedChangeStates = []string{"private", "remove-private", "wip",
	"ready"}

// Commit message
var Message string

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
	rootCmd.Version = "v1.0.0-rc"

	rootCmd.PersistentFlags().StringVarP(&Branch, "branch", "b", "",
		"Target branch name")
	rootCmd.PersistentFlags().StringVarP(&State, "state", "s", "",
		"Change state: private, remove-private, wip or ready")
	rootCmd.PersistentFlags().StringVarP(&Message, "message", "m", "",
		"Commit message")
}
