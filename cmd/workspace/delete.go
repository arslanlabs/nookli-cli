package workspace

import (
	dbw "nookli/db/workspace"

	"github.com/spf13/cobra"
)

var (
	deleteID      int
	deleteConfirm bool
)

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a workspace",
	Run: func(cmd *cobra.Command, args []string) {
		if deleteID == 0 || !deleteConfirm {
			cmd.PrintErrln("Error: --id and --yes are required")
			return
		}
		if err := dbw.Delete(deleteID); err != nil {
			cmd.PrintErrln("Error deleting workspace:", err)
			return
		}
		cmd.Println("Workspace deleted:", deleteID)
	},
}

func init() {
	deleteCmd.Flags().IntVarP(&deleteID, "id", "i", 0, "Workspace ID (required)")
	deleteCmd.Flags().BoolVarP(&deleteConfirm, "yes", "y", false, "Confirm deletion")
}
