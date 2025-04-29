package workspace_test

import (
	"bytes"
	"os"
	"strings"
	"testing"

	"nookli/cmd"
	db "nookli/db"
)

// runCLI drives the workspace subcommands, closing the DB before and after each run.
func runCLI(t *testing.T, args ...string) (string, error) {
	t.Helper()

	// 1) Close any open DB handle
	if db.DB != nil {
		db.DB.Close()
		db.DB = nil
	}

	// 2) Prepare and run the workspace command
	root := cmd.GetRootCmd()
	buf := &bytes.Buffer{}
	root.SetOut(buf)
	root.SetErr(buf)
	// Prefix with "workspace"
	root.SetArgs(append([]string{"workspace"}, args...))
	_, err := root.ExecuteC()

	// 3) Close the DB again so the TempDir cleanup can delete the file
	if db.DB != nil {
		db.DB.Close()
		db.DB = nil
	}

	return buf.String(), err
}

func TestCRUD(t *testing.T) {
	tmp := t.TempDir()
	cwd, _ := os.Getwd()
	defer os.Chdir(cwd)
	os.Chdir(tmp)

	// CREATE
	out, _ := runCLI(t, "create", "--name", "W", "--description", "D")
	if !strings.Contains(out, "Workspace created: W") {
		t.Fatalf("create failed, got %q", out)
	}

	// LIST
	out, _ = runCLI(t, "list")
	if !strings.Contains(out, "1: W — D") {
		t.Fatalf("list failed, got %q", out)
	}

	// SHOW
	out, _ = runCLI(t, "show", "--id", "1")
	if !strings.Contains(out, "1: W — D") {
		t.Fatalf("show failed, got %q", out)
	}

	// UPDATE
	out, _ = runCLI(t, "update", "--id", "1", "--name", "X")
	if !strings.Contains(out, "Workspace updated: 1") {
		t.Fatalf("update failed, got %q", out)
	}

	// DELETE (missing --yes)
	out, _ = runCLI(t, "delete", "--id", "1")
	if !strings.Contains(out, "Error: --id and --yes are required") {
		t.Fatalf("delete flag error missing, got %q", out)
	}

	// DELETE (with --yes)
	out, _ = runCLI(t, "delete", "--id", "1", "--yes")
	if !strings.Contains(out, "Workspace deleted: 1") {
		t.Fatalf("delete failed, got %q", out)
	}
}
