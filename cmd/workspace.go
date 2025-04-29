// cmd/workspace.go
package cmd

import (
	db "nookli/db"

	"github.com/spf13/cobra"
)

var (
	wsName string
	wsDesc string
)

var workspaceCmd = &cobra.Command{
	Use:   "workspace",
	Short: "Manage workspaces",
	Long:  "Create and list workspaces in your local Nookli database.",
}

var workspaceCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new workspace",
	Long:  "Creates a new workspace. --name is required, --description is optional.",
	Run: func(cmd *cobra.Command, args []string) {
		if wsName == "" {
			cmd.PrintErrln("Error: --name is required")
			return
		}
		if err := db.CreateWorkspace(wsName, wsDesc); err != nil {
			cmd.PrintErrln("Error creating workspace:", err)
			return
		}
		cmd.Println("Workspace created:", wsName)
	},
}

var workspaceListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all workspaces",
	Long:  "Lists every workspace stored in your local database.",
	Run: func(cmd *cobra.Command, args []string) {
		wks, err := db.ListWorkspaces()
		if err != nil {
			cmd.PrintErrln("Error listing workspaces:", err)
			return
		}
		cmd.Println("Workspaces:")
		for _, w := range wks {
			cmd.Printf("  %d: %s — %s\n", w.ID, w.Name, w.Description)
		}
	},
}

func init() {
	rootCmd.AddCommand(workspaceCmd)

	// workspace create
	workspaceCmd.AddCommand(workspaceCreateCmd)
	workspaceCreateCmd.Flags().StringVarP(&wsName, "name", "n", "", "Workspace name (required)")
	workspaceCreateCmd.Flags().StringVarP(&wsDesc, "description", "d", "", "Workspace description")

	// workspace list
	workspaceCmd.AddCommand(workspaceListCmd)
}
