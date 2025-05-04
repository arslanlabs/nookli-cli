package workspace

import (
	"context"

	svcw "nookli/pkg/service/workspace"

	"github.com/spf13/cobra"
)

var (
	updateID   int
	updateName string
	updateDesc string
)

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update a workspace",
	Run: func(cmd *cobra.Command, args []string) {
		if updateID == 0 {
			cmd.PrintErrln("Error: --id is required")
			return
		}
		if updateName == "" && updateDesc == "" {
			cmd.PrintErrln("Error: --name or --description is required")
			return
		}
		svc := svcw.NewService()
		if err := svc.Update(context.Background(), updateID, updateName, updateDesc); err != nil {
			cmd.PrintErrln("Error updating workspace:", err)
			return
		}
		cmd.Println("Workspace updated:", updateID)
	},
}

func init() {
	updateCmd.Flags().IntVarP(&updateID, "id", "i", 0, "Workspace ID (required)")
	updateCmd.Flags().StringVarP(&updateName, "name", "n", "", "New name")
	updateCmd.Flags().StringVarP(&updateDesc, "description", "d", "", "New description")
}
