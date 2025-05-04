package cmd

import (
	"fmt"
	"os"

	workspacecmd "nookli/cmd/workspace"
	"nookli/db"

	"github.com/spf13/cobra"
)

var (
	version = "dev" // populated via -ldflags
	commit  = "none"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the CLI version",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Printf("nookli version %s (commit %s)\n", version, commit)
	},
}

// rootCmd is the base command.
var rootCmd = &cobra.Command{
	Use:   "nookli",
	Short: "Nookli CLI — your Knowledge OS command line",
	Long: `Nookli is a developer tool to manage workspaces, stacks, elements
and dynamic learning paths right from your terminal.`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		// Initialize the DB (tables etc.)
		db.InitDB()
	},
}

// RootCmd exported for tests
var RootCmd = rootCmd

func init() {
	rootCmd.AddCommand(versionCmd)

	// Register workspace commands
	rootCmd.AddCommand(workspacecmd.Cmd)
	// future: rootCmd.AddCommand(stackcmd.Cmd)
	// future: rootCmd.AddCommand(elementcmd.Cmd)
	// future: rootCmd.AddCommand(blockcmd.Cmd)
}

// Execute runs the root command.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// GetRootCmd returns the root for tests.
func GetRootCmd() *cobra.Command {
	return rootCmd
}
