// cmd/cli_test.go
package cmd

import (
	"bytes"
	"testing"

	db "nookli/db"
)

// runCLI runs the CLI in the current directory, resetting state
// and closing the DB to avoid file‐lock issues.
func runCLI(t *testing.T, args ...string) (string, error) {
	t.Helper()

	// 1) Close any open DB from prior run
	if db.DB != nil {
		db.DB.Close()
		db.DB = nil
	}

	// 2) Reset all package‐level flags
	wsName = ""
	wsDesc = ""
	stackName = ""
	stackWSID = 0
	blockName = ""
	blockContent = ""
	blockStackID = 0
	elemName = ""
	elemDesc = ""

	// 3) Execute the command
	root := GetRootCmd()
	buf := &bytes.Buffer{}
	root.SetOut(buf)
	root.SetErr(buf)
	root.SetArgs(args)

	_, err := root.ExecuteC()

	// 4) Close the DB again so TempDir cleanup succeeds
	if db.DB != nil {
		db.DB.Close()
		db.DB = nil
	}

	return buf.String(), err
}

// contains is a small helper to check substrings.
func contains(output, substr string) bool {
	return bytes.Contains([]byte(output), []byte(substr))
}
