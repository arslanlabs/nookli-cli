package workspace

import (
	dbw "nookli/db/workspace"

	"github.com/spf13/cobra"
)

var (
	createName string
	createDesc string
)

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new workspace",
	Run: func(cmd *cobra.Command, args []string) {
		if createName == "" {
			cmd.PrintErrln("Error: --name is required")
			return
		}
		if err := dbw.Create(createName, createDesc); err != nil {
			cmd.PrintErrln("Error creating workspace:", err)
			return
		}
		cmd.Println("Workspace created:", createName)
	},
}

func init() {
	createCmd.Flags().StringVarP(&createName, "name", "n", "", "Workspace name (required)")
	createCmd.Flags().StringVarP(&createDesc, "description", "d", "", "Workspace description")
}
