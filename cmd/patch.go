package cmd

import (
	"github.com/spf13/cobra"

	"github.com/aruncveli/gerritr/pkg/action"
)

var patchCmd = &cobra.Command{
	Use:   "patch",
	Short: "Patch the current review with the staged files",
	Run: func(cmd *cobra.Command, args []string) {
		action.Patch(Branch, State, Message)
	},
}

func init() {
	rootCmd.AddCommand(patchCmd)
}
