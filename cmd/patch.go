package cmd

import (
	"github.com/spf13/cobra"

	"gerritr/pkg/action"
)

var patchCmd = &cobra.Command{
	Use:   "patch",
	Short: "Patch the current review with the added/staged files",
	Run: func(cmd *cobra.Command, args []string) {
		action.Patch(Branch, State, Message)
	},
}

func init() {
	rootCmd.AddCommand(patchCmd)

	// patchCmd.Flags().StringVarP(&message, "message", "m", "", "Commit message")
}
