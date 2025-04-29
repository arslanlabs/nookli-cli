// cmd/workspace_test.go
package cmd

import (
	"os"
	"testing"
)

func TestWorkspaceCommands(t *testing.T) {
	// Use one temp dir for both create and list
	tmp := t.TempDir()
	cwd, _ := os.Getwd()
	defer os.Chdir(cwd)
	os.Chdir(tmp)

	// 1) Missing --name should error
	out, _ := runCLI(t, "workspace", "create")
	if !contains(out, "Error: --name is required") {
		t.Errorf("expected missing-name error, got %q", out)
	}

	// 2) Create workspace
	out, err := runCLI(t,
		"workspace", "create",
		"--name", "Foo",
		"--description", "Bar",
	)
	if err != nil {
		t.Fatalf("workspace create failed: %v", err)
	}
	if !contains(out, "Workspace created: Foo") {
		t.Errorf("unexpected create output: %q", out)
	}

	// 3) List workspaces
	out, err = runCLI(t, "workspace", "list")
	if err != nil {
		t.Fatalf("workspace list failed: %v", err)
	}
	if !contains(out, "1: Foo — Bar") {
		t.Errorf("unexpected list output: %q", out)
	}
}
