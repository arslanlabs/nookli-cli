package workspace

import (
	"encoding/json"
	"fmt"
	"text/tabwriter"

	dbw "nookli/db/workspace"

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
		wks, err := dbw.List()
		if err != nil {
			cmd.PrintErrln("Error listing workspaces:", err)
			return
		}
		out := cmd.OutOrStdout()

		// JSON output
		if listOutput == "json" {
			enc := json.NewEncoder(out)
			enc.SetIndent("", "  ")
			if err := enc.Encode(wks); err != nil {
				cmd.PrintErrln("Error encoding JSON:", err)
			}
			return
		}

		// Verbose tabular output
		if listVerbose {
			tw := tabwriter.NewWriter(out, 0, 4, 2, ' ', 0)
			// header
			fmt.Fprintln(tw, "ID\tName\tDescription\tCreatedAt")
			for _, w := range wks {
				fmt.Fprintf(tw, "%d\t%s\t%s\t%s\n",
					w.ID, w.Name, w.Description, w.CreatedAt)
			}
			tw.Flush()
			return
		}

		// Simple text output
		for _, w := range wks {
			fmt.Fprintf(out, "%d: %s — %s\n", w.ID, w.Name, w.Description)
		}
	},
}

func init() {
	listCmd.Flags().BoolVarP(&listVerbose, "verbose", "v", false, "Show all columns")
	listCmd.Flags().StringVarP(&listOutput, "output", "o", "text", "Output (text|json)")
}
