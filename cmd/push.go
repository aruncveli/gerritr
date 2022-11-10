package cmd

import (
	"github.com/spf13/cobra"

	"github.com/aruncveli/gerritr/pkg/action"
)

// Space separated list of reviewer email IDs or aliases
var reviewers []string

var pushCmd = &cobra.Command{
	Use:   "push",
	Short: "Push the latest commit and add reviewers",
	Run: func(cmd *cobra.Command, args []string) {
		action.Push(Branch, State, Message, reviewers)
	},
}

func init() {
	rootCmd.AddCommand(pushCmd)

	pushCmd.Flags().StringArrayVarP(&reviewers, "reviewers", "r", nil,
		"Space separated list of reviewer email IDs or aliases")
}
