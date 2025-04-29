package cmd

import (
	"bytes"
	"testing"

	db "nookli/db"
)

// runCLI runs the CLI in the current directory, resetting all cmd-package flags
// and closing the DB to avoid file-lock issues.
func runCLI(t *testing.T, args ...string) (string, error) {
	t.Helper()

	// 1) Close any open DB from a prior run
	if db.DB != nil {
		db.DB.Close()
		db.DB = nil
	}

	// 2) Reset all package-level flags for cmd/*
	stackName = ""
	stackWSID = 0

	elemName = ""
	elemDesc = ""

	blockName = ""
	blockContent = ""
	blockStackID = 0

	// 3) Execute the command
	root := GetRootCmd()
	buf := &bytes.Buffer{}
	root.SetOut(buf)
	root.SetErr(buf)
	root.SetArgs(args)
	_, err := root.ExecuteC()

	// 4) Close the DB again so TempDir cleanup in tests succeeds
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
