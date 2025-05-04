package workspace

import (
	"context"

	svcw "nookli/pkg/service/workspace"

	"github.com/spf13/cobra"
)

var (
	createName string
	createDesc string
)

// Setup "create" command.
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new workspace",
	Run: func(cmd *cobra.Command, args []string) {
		if createName == "" {
			cmd.PrintErrln("Error: --name is required")
			return
		}
		svc := svcw.NewService()
		w, err := svc.Create(context.Background(), createName, createDesc)
		if err != nil {
			cmd.PrintErrln("Error creating workspace:", err)
			return
		}
		cmd.Println("Workspace created:", w.Name)
	},
}

func init() {
	createCmd.Flags().StringVarP(&createName, "name", "n", "", "Workspace name (required)")
	createCmd.Flags().StringVarP(&createDesc, "description", "d", "", "Workspace description")
}
