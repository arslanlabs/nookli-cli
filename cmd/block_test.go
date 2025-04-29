// cmd/block_test.go
package cmd

import (
	"os"
	"testing"
)

func TestBlockCommands(t *testing.T) {
	// 1) Switch into a single temp dir for all CLI calls in this test
	tmp := t.TempDir()
	cwd, _ := os.Getwd()
	defer os.Chdir(cwd)
	os.Chdir(tmp)

	// 2) Missing flags
	out, _ := runCLI(t, "block", "create")
	if !contains(out, "Error: --stack-id and --name are required") {
		t.Errorf("expected missing-flag error, got %q", out)
	}

	// 3) Prep workspace & stack in the same dir
	if _, err := runCLI(t, "workspace", "create", "--name", "W", "--description", ""); err != nil {
		t.Fatalf("workspace create failed: %v", err)
	}
	if _, err := runCLI(t, "stack", "create", "--workspace-id", "1", "--name", "S"); err != nil {
		t.Fatalf("stack create failed: %v", err)
	}

	// 4) Create block
	out, err := runCLI(t, "block", "create", "--stack-id", "1", "--name", "B1", "--content", "C1")
	if err != nil {
		t.Fatalf("block create failed: %v", err)
	}
	if !contains(out, "Block created: B1") {
		t.Errorf("unexpected create output: %q", out)
	}

	// 5) List blocks (same DB file, same dir)
	out, err = runCLI(t, "block", "list", "--stack-id", "1")
	if err != nil {
		t.Fatalf("block list failed: %v", err)
	}
	if !contains(out, "1: B1 — C1") {
		t.Errorf("expected block in list, got %q", out)
	}
}
