package cmd

import (
	"fmt"
	"os"

	"nookli/db"

	workspacecmd "nookli/cmd/workspace"

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

// private root
var rootCmd = &cobra.Command{
	Use:   "nookli",
	Short: "Nookli CLI — your Knowledge OS command line",
	Long: `Nookli is a developer tool to manage workspaces, stacks, elements,
and dynamic learning paths right from your terminal.`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		db.InitDB()
	},
}

// **THIS** lets tests refer to it as cmd.RootCmd
var RootCmd = rootCmd

func init() {
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(workspacecmd.Cmd)
	// rootCmd.AddCommand(stackcmd.Cmd)    // when you migrate stack
	// rootCmd.AddCommand(elementcmd.Cmd)
	// rootCmd.AddCommand(blockcmd.Cmd)
}

// Execute kicks things off
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// For tests you can also use GetRootCmd()
func GetRootCmd() *cobra.Command {
	return rootCmd
}
