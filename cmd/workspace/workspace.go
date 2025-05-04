package workspace

import (
	"github.com/spf13/cobra"
)

// Cmd is the entry‐point for all workspace subcommands.
var Cmd = &cobra.Command{
	Use:   "workspace",
	Short: "Manage workspaces",
	Long:  "Create, list, show, update and delete workspaces.",
}

func init() {
	Cmd.AddCommand(
		createCmd,
		listCmd,
		showCmd,
		updateCmd,
		deleteCmd,
	)
}
