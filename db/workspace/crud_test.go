package workspace

import (
	"os"
	"testing"

	db "nookli/db"
)

// runInTempDir ensures each test has its own working directory and fresh DB file.
func runInTempDir(t *testing.T, fn func()) {
	t.Helper()
	orig, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	tmp := t.TempDir()
	if err := os.Chdir(tmp); err != nil {
		t.Fatal(err)
	}
	defer func() {
		if db.DB != nil {
			db.DB.Close()
			db.DB = nil
		}
		os.Chdir(orig)
	}()

	fn()
}

func TestWorkspaceCRUD(t *testing.T) {
	runInTempDir(t, func() {
		// Initialize the DB (creates tables)
		db.InitDB()

		// Create + List
		if err := Create("A", "B"); err != nil {
			t.Fatal(err)
		}
		wks, err := List()
		if err != nil {
			t.Fatal(err)
		}
		if len(wks) != 1 || wks[0].Name != "A" {
			t.Fatalf("unexpected list: %+v", wks)
		}

		// Get
		w, err := Get(wks[0].ID)
		if err != nil {
			t.Fatal(err)
		}
		if w.Description != "B" {
			t.Fatalf("unexpected get: %+v", w)
		}

		// Update & re-Get
		if err := Update(w.ID, "X", "Y"); err != nil {
			t.Fatal(err)
		}
		w2, err := Get(w.ID)
		if err != nil {
			t.Fatal(err)
		}
		if w2.Name != "X" || w2.Description != "Y" {
			t.Fatalf("unexpected update: %+v", w2)
		}

		// Delete & ensure gone
		if err := Delete(w.ID); err != nil {
			t.Fatal(err)
		}
		if _, err := Get(w.ID); err == nil {
			t.Fatal("expected error on get after delete")
		}
	})
}
