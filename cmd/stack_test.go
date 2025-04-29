// cmd/stack_test.go
package cmd

import (
	"os"
	"testing"
)

func TestStackCommands(t *testing.T) {
	// use one temp dir so our sqlite file is shared across calls
	tmp := t.TempDir()
	cwd, _ := os.Getwd()
	defer os.Chdir(cwd)
	os.Chdir(tmp)

	// 1) Missing flags
	out, _ := runCLI(t, "stack", "create")
	if !contains(out, "Error: --workspace-id and --name are required") {
		t.Errorf("expected missing-flags error, got %q", out)
	}

	// 2) Create a workspace so that workspace-id=1 exists
	if _, err := runCLI(t, "workspace", "create", "--name", "W", "--description", ""); err != nil {
		t.Fatalf("workspace create failed: %v", err)
	}

	// 3) Create the stack
	out, err := runCLI(t,
		"stack", "create",
		"--workspace-id", "1",
		"--name", "S1",
	)
	if err != nil {
		t.Fatalf("stack create failed: %v", err)
	}
	if !contains(out, "Stack created: S1") {
		t.Errorf("unexpected create output: %q", out)
	}

	// 4) List stacks
	out, err = runCLI(t, "stack", "list", "--workspace-id", "1")
	if err != nil {
		t.Fatalf("stack list failed: %v", err)
	}
	if !contains(out, "1: S1") {
		t.Errorf("unexpected list output: %q", out)
	}
}
