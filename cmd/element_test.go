// cmd/element_test.go
package cmd

import (
	"os"
	"testing"
)

func TestElementCommands(t *testing.T) {
	tmp := t.TempDir()
	cwd, _ := os.Getwd()
	defer os.Chdir(cwd)
	os.Chdir(tmp)

	// Missing name
	out, _ := runCLI(t, "element", "create")
	if !contains(out, "Error: --name is required") {
		t.Errorf("expected missing-name error, got %q", out)
	}

	// Create element
	out, err := runCLI(t, "element", "create", "--name", "E1", "--description", "D1")
	if err != nil {
		t.Fatalf("element create failed: %v", err)
	}
	if !contains(out, "Element created: E1") {
		t.Errorf("unexpected create output: %q", out)
	}

	// List elements
	out, err = runCLI(t, "element", "list")
	if err != nil {
		t.Fatalf("element list failed: %v", err)
	}
	if !contains(out, "1: E1 — D1") {
		t.Errorf("expected element in list, got %q", out)
	}
}
