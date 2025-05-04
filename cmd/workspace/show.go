package workspace

import (
	"context"

	svcw "nookli/pkg/service/workspace"

	"github.com/spf13/cobra"
)

var showID int

var showCmd = &cobra.Command{
	Use:   "show",
	Short: "Show one workspace",
	Run: func(cmd *cobra.Command, args []string) {
		if showID == 0 {
			cmd.PrintErrln("Error: --id is required")
			return
		}
		svc := svcw.NewService()
		w, err := svc.Get(context.Background(), showID)
		if err != nil {
			cmd.PrintErrln("Error fetching workspace:", err)
			return
		}
		cmd.Printf("%d: %s — %s\n", w.ID, w.Name, w.Description)
	},
}

func init() {
	showCmd.Flags().IntVarP(&showID, "id", "i", 0, "Workspace ID (required)")
}
