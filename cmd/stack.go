// cmd/stack.go
package cmd

import (
	db "nookli/db"

	"github.com/spf13/cobra"
)

var (
	stackName string
	stackWSID int
)

var stackCmd = &cobra.Command{
	Use:   "stack",
	Short: "Manage stacks",
	Long:  "Create and list stacks under a given workspace.",
}

var stackCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new stack",
	Long:  "Creates a new stack under --workspace-id. Both flags are required.",
	Run: func(cmd *cobra.Command, args []string) {
		if stackName == "" || stackWSID == 0 {
			cmd.PrintErrln("Error: --workspace-id and --name are required")
			return
		}
		if err := db.CreateStack(stackName, stackWSID); err != nil {
			cmd.PrintErrln("Error creating stack:", err)
			return
		}
		cmd.Println("Stack created:", stackName)
	},
}

var stackListCmd = &cobra.Command{
	Use:   "list",
	Short: "List stacks in a workspace",
	Long:  "Lists all stacks belonging to --workspace-id.",
	Run: func(cmd *cobra.Command, args []string) {
		if stackWSID == 0 {
			cmd.PrintErrln("Error: --workspace-id is required")
			return
		}
		stacks, err := db.ListStacks(stackWSID)
		if err != nil {
			cmd.PrintErrln("Error listing stacks:", err)
			return
		}
		cmd.Println("Stacks:")
		for _, s := range stacks {
			cmd.Printf("  %d: %s (Workspace %d)\n", s.ID, s.Name, s.WorkspaceID)
		}
	},
}

func init() {
	rootCmd.AddCommand(stackCmd)

	// stack create flags
	stackCmd.AddCommand(stackCreateCmd)
	stackCreateCmd.Flags().StringVarP(&stackName, "name", "n", "", "Stack name (required)")
	stackCreateCmd.Flags().IntVarP(&stackWSID, "workspace-id", "w", 0, "Workspace ID (required)")

	// stack list flags
	stackCmd.AddCommand(stackListCmd)
	stackListCmd.Flags().IntVarP(&stackWSID, "workspace-id", "w", 0, "Workspace ID (required)")
}
