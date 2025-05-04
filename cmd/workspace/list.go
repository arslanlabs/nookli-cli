package workspace

import (
	"context"
	"encoding/json"
	"fmt"
	"text/tabwriter"

	svcw "nookli/pkg/service/workspace"

	"github.com/spf13/cobra"
)

var (
	listVerbose bool
	listOutput  string
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all workspaces",
	Run: func(cmd *cobra.Command, args []string) {
		svc := svcw.NewService()
		wks, err := svc.List(context.Background())
		if err != nil {
			cmd.PrintErrln("Error listing workspaces:", err)
			return
		}
		out := cmd.OutOrStdout()

		if listOutput == "json" {
			enc := json.NewEncoder(out)
			enc.SetIndent("", "  ")
			enc.Encode(wks)
			return
		}
		if listVerbose {
			tw := tabwriter.NewWriter(out, 0, 4, 2, ' ', 0)
			fmt.Fprintln(tw, "ID\tName\tDescription\tCreatedAt")
			for _, w := range wks {
				fmt.Fprintf(tw, "%d\t%s\t%s\t%s\n",
					w.ID, w.Name, w.Description, w.CreatedAt)
			}
			tw.Flush()
			return
		}
		for _, w := range wks {
			fmt.Fprintf(out, "%d: %s — %s\n", w.ID, w.Name, w.Description)
		}
	},
}

func init() {
	listCmd.Flags().BoolVarP(&listVerbose, "verbose", "v", false, "Show all columns")
	listCmd.Flags().StringVarP(&listOutput, "output", "o", "text", "Output (text|json)")
}
