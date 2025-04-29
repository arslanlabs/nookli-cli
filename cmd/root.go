// cmd/root.go
package cmd

import (
	"fmt"
	"os"

	"nookli/db"

	"github.com/spf13/cobra"
)

var (
	// Populated via -ldflags
	version = "dev"
	commit  = "none"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the CLI version",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Printf("nookli version %s (commit %s)\n", version, commit)
	},
}

var rootCmd = &cobra.Command{
	Use:   "nookli",
	Short: "Nookli CLI — your Knowledge OS command line",
	Long: `Nookli is a developer tool to manage workspaces, stacks, elements,
and dynamic learning paths right from your terminal.`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		// Ensure the DB is ready before any subcommand runs.
		db.InitDB()
	},
}

// Execute kicks off the Cobra command parsing.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	// Here you can define global flags, e.g.:
	// rootCmd.PersistentFlags().BoolP("verbose", "v", false, "verbose output")
	rootCmd.AddCommand(versionCmd)

}

// At the bottom of cmd/root.go
func GetRootCmd() *cobra.Command {
	return rootCmd
}
