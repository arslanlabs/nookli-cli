// db/sqlite_test.go
package db

import (
	"os"
	"testing"
)

// runInTempDir switches to a temp dir so each test gets a fresh DB file,
// and ensures the DB handle is closed so Windows can delete the file.
func runInTempDir(t *testing.T, fn func()) {
	t.Helper()

	// Remember original working directory
	orig, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}

	// Create and switch to a new temp directory
	tmp := t.TempDir()
	if err := os.Chdir(tmp); err != nil {
		t.Fatal(err)
	}

	// When the test function completes:
	defer func() {
		// Close the DB handle so the file lock is released
		if DB != nil {
			DB.Close()
			DB = nil
		}
		// Switch back to the original directory
		os.Chdir(orig)
	}()

	// Run the test’s logic
	fn()
}

func TestWorkspaceCRUD(t *testing.T) {
	runInTempDir(t, func() {
		InitDB()

		// Initially empty
		wks, err := ListWorkspaces()
		if err != nil {
			t.Fatal(err)
		}
		if len(wks) != 0 {
			t.Fatalf("expected 0 workspaces, got %d", len(wks))
		}

		// Create one
		if err := CreateWorkspace("X", "Y"); err != nil {
			t.Fatal(err)
		}

		// Verify it’s listed
		wks, err = ListWorkspaces()
		if err != nil {
			t.Fatal(err)
		}
		if len(wks) != 1 {
			t.Fatalf("expected 1 workspace, got %d", len(wks))
		}
		if got := wks[0]; got.Name != "X" || got.Description != "Y" {
			t.Errorf("unexpected workspace: %+v", got)
		}
	})
}

func TestStackCRUD(t *testing.T) {
	runInTempDir(t, func() {
		InitDB()
		CreateWorkspace("W", "")
		wks, _ := ListWorkspaces()
		wid := wks[0].ID

		// No stacks yet
		st, err := ListStacks(wid)
		if err != nil {
			t.Fatal(err)
		}
		if len(st) != 0 {
			t.Fatalf("expected 0 stacks, got %d", len(st))
		}

		// Create and verify
		if err := CreateStack("S", wid); err != nil {
			t.Fatal(err)
		}
		st, err = ListStacks(wid)
		if err != nil {
			t.Fatal(err)
		}
		if len(st) != 1 || st[0].Name != "S" {
			t.Fatalf("unexpected stacks: %+v", st)
		}
	})
}

func TestElementCRUD(t *testing.T) {
	runInTempDir(t, func() {
		InitDB()

		// Empty at first
		els, _ := ListElements()
		if len(els) != 0 {
			t.Fatalf("expected 0 elements, got %d", len(els))
		}

		// Create and verify
		if err := CreateElement("E", "D"); err != nil {
			t.Fatal(err)
		}
		els, err := ListElements()
		if err != nil {
			t.Fatal(err)
		}
		if len(els) != 1 || els[0].Name != "E" {
			t.Fatalf("unexpected elements: %+v", els)
		}
	})
}

func TestBlockCRUD(t *testing.T) {
	runInTempDir(t, func() {
		InitDB()
		CreateWorkspace("W", "")
		wks, _ := ListWorkspaces()
		wid := wks[0].ID
		CreateStack("S", wid)
		st, _ := ListStacks(wid)
		sid := st[0].ID

		// Empty initially
		bl, _ := ListBlocks(sid)
		if len(bl) != 0 {
			t.Fatalf("expected 0 blocks, got %d", len(bl))
		}

		// Create and verify
		if err := CreateBlock("B", "C", sid); err != nil {
			t.Fatal(err)
		}
		bl, err := ListBlocks(sid)
		if err != nil {
			t.Fatal(err)
		}
		if len(bl) != 1 || bl[0].Name != "B" || bl[0].Content != "C" {
			t.Fatalf("unexpected blocks: %+v", bl)
		}
	})
}
