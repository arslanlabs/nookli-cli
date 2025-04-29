// integration_test.go
package main

import (
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
)

// buildCLI builds your nookli binary in the project root.
func buildCLI(t *testing.T) string {
	t.Helper()
	bin := "nookli"
	if runtime.GOOS == "windows" {
		bin += ".exe"
	}
	// remove any stale binary
	os.Remove(bin)

	cmd := exec.Command("go", "build", "-o", bin)
	cmd.Env = os.Environ()
	if out, err := cmd.CombinedOutput(); err != nil {
		t.Fatalf("go build failed: %v\n%s", err, out)
	}
	return filepath.Join(".", bin)
}

// run runs `bin args...` in dir, fatally fails on error, and returns stdout+stderr.
func run(t *testing.T, bin, dir string, args ...string) string {
	t.Helper()
	cmd := exec.Command(bin, args...)
	cmd.Dir = dir
	out, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("`%s %s` failed: %v\n%s",
			bin, strings.Join(args, " "), err, out)
	}
	return string(out)
}

func TestIntegration_CreateAndListAll(t *testing.T) {
	bin := buildCLI(t)

	// 1) Single temp dir for all commands (one DB file)
	tmp := t.TempDir()

	// Workspace
	t.Run("Workspace", func(t *testing.T) {
		out := run(t, bin, tmp, "workspace", "create", "--name", "W1", "--description", "D1")
		if !strings.Contains(out, "Workspace created: W1") {
			t.Fatalf("unexpected create output: %s", out)
		}
		out = run(t, bin, tmp, "workspace", "list")
		if !strings.Contains(out, "1: W1 — D1") {
			t.Fatalf("unexpected list output: %s", out)
		}
	})

	// Stack
	t.Run("Stack", func(t *testing.T) {
		out := run(t, bin, tmp, "stack", "create", "--workspace-id", "1", "--name", "S1")
		if !strings.Contains(out, "Stack created: S1") {
			t.Fatalf("unexpected create output: %s", out)
		}
		out = run(t, bin, tmp, "stack", "list", "--workspace-id", "1")
		if !strings.Contains(out, "1: S1") {
			t.Fatalf("unexpected list output: %s", out)
		}
	})

	// Element
	t.Run("Element", func(t *testing.T) {
		out := run(t, bin, tmp, "element", "create", "--name", "E1", "--description", "D1")
		if !strings.Contains(out, "Element created: E1") {
			t.Fatalf("unexpected create output: %s", out)
		}
		out = run(t, bin, tmp, "element", "list")
		if !strings.Contains(out, "1: E1 — D1") {
			t.Fatalf("unexpected list output: %s", out)
		}
	})

	// Block
	t.Run("Block", func(t *testing.T) {
		// Need stack 1 already created above
		out := run(t, bin, tmp,
			"block", "create",
			"--stack-id", "1",
			"--name", "B1",
			"--content", "C1",
		)
		if !strings.Contains(out, "Block created: B1") {
			t.Fatalf("unexpected create output: %s", out)
		}
		out = run(t, bin, tmp, "block", "list", "--stack-id", "1")
		if !strings.Contains(out, "1: B1 — C1") {
			t.Fatalf("unexpected list output: %s", out)
		}
	})
}
