// cmd/element.go
package cmd

import (
	db "nookli/db"

	"github.com/spf13/cobra"
)

var (
	elemName string
	elemDesc string
)

var elementCmd = &cobra.Command{
	Use:   "element",
	Short: "Manage elements",
	Long:  "Create and list elements that wrap content as metadata.",
}

var elementCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new element",
	Long:  "Creates a new element. --name is required, --description optional.",
	Run: func(cmd *cobra.Command, args []string) {
		if elemName == "" {
			cmd.PrintErrln("Error: --name is required")
			return
		}
		if err := db.CreateElement(elemName, elemDesc); err != nil {
			cmd.PrintErrln("Error creating element:", err)
			return
		}
		cmd.Println("Element created:", elemName)
	},
}

var elementListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all elements",
	Long:  "Lists every element stored in your local database.",
	Run: func(cmd *cobra.Command, args []string) {
		elems, err := db.ListElements()
		if err != nil {
			cmd.PrintErrln("Error listing elements:", err)
			return
		}
		cmd.Println("Elements:")
		for _, e := range elems {
			cmd.Printf("  %d: %s — %s\n", e.ID, e.Name, e.Description)
		}
	},
}

func init() {
	rootCmd.AddCommand(elementCmd)

	elementCmd.AddCommand(elementCreateCmd)
	elementCreateCmd.Flags().StringVarP(&elemName, "name", "n", "", "Element name (required)")
	elementCreateCmd.Flags().StringVarP(&elemDesc, "description", "d", "", "Element description")

	elementCmd.AddCommand(elementListCmd)
}
