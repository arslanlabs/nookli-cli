// cmd/block.go
package cmd

import (
	db "nookli/db"

	"github.com/spf13/cobra"
)

var (
	blockName    string
	blockContent string
	blockStackID int
)

var blockCmd = &cobra.Command{
	Use:   "block",
	Short: "Manage blocks",
	Long:  "Create and list blocks (content units) within a stack.",
}

var blockCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new block",
	Long:  "Creates a block under --stack-id with --name and optional --content.",
	Run: func(cmd *cobra.Command, args []string) {
		if blockName == "" || blockStackID == 0 {
			cmd.PrintErrln("Error: --stack-id and --name are required")
			return
		}
		if err := db.CreateBlock(blockName, blockContent, blockStackID); err != nil {
			cmd.PrintErrln("Error creating block:", err)
			return
		}
		cmd.Println("Block created:", blockName)
	},
}

var blockListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all blocks in a stack",
	Long:  "Lists every block belonging to the given --stack-id.",
	Run: func(cmd *cobra.Command, args []string) {
		if blockStackID == 0 {
			cmd.PrintErrln("Error: --stack-id is required")
			return
		}
		blks, err := db.ListBlocks(blockStackID)
		if err != nil {
			cmd.PrintErrln("Error listing blocks:", err)
			return
		}
		cmd.Println("Blocks:")
		for _, b := range blks {
			cmd.Printf("  %d: %s — %s (Stack %d)\n",
				b.ID, b.Name, b.Content, b.StackID)
		}
	},
}

func init() {
	rootCmd.AddCommand(blockCmd)

	// block create
	blockCmd.AddCommand(blockCreateCmd)
	blockCreateCmd.Flags().StringVarP(&blockName, "name", "n", "", "Block name (required)")
	blockCreateCmd.Flags().StringVarP(&blockContent, "content", "c", "", "Block content")
	blockCreateCmd.Flags().IntVarP(&blockStackID, "stack-id", "s", 0, "Stack ID (required)")

	// block list
	blockCmd.AddCommand(blockListCmd)
	blockListCmd.Flags().IntVarP(&blockStackID, "stack-id", "s", 0, "Stack ID (required)")
}
