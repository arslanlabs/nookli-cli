// db/sqlite.go
package db

import (
	"database/sql"
	"log"

	// cgo driver for SQLite
	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

// InitDB opens/creates the DB file and ensures all tables exist.
func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", "./nookli.db")
	if err != nil {
		log.Fatal("opening database:", err)
	}
	createTables()
}

func createTables() {
	// Workspaces table (with created_at)
	if _, err := DB.Exec(`
        CREATE TABLE IF NOT EXISTS workspaces (
            id          INTEGER PRIMARY KEY AUTOINCREMENT,
            name        TEXT NOT NULL,
            description TEXT,
            created_at  DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
        );
    `); err != nil {
		log.Fatal("creating workspaces table:", err)
	}

	// Stacks table (linked to workspace)
	if _, err := DB.Exec(`
        CREATE TABLE IF NOT EXISTS stacks (
            id            INTEGER PRIMARY KEY AUTOINCREMENT,
            name          TEXT NOT NULL,
            workspace_id  INTEGER NOT NULL
        );
    `); err != nil {
		log.Fatal("creating stacks table:", err)
	}

	// Elements table
	if _, err := DB.Exec(`
        CREATE TABLE IF NOT EXISTS elements (
            id          INTEGER PRIMARY KEY AUTOINCREMENT,
            name        TEXT NOT NULL,
            description TEXT
        );
    `); err != nil {
		log.Fatal("creating elements table:", err)
	}

	// Blocks table: each block belongs to a stack, has content
	if _, err := DB.Exec(`
        CREATE TABLE IF NOT EXISTS blocks (
            id       INTEGER PRIMARY KEY AUTOINCREMENT,
            name     TEXT NOT NULL,
            content  TEXT,
            stack_id INTEGER NOT NULL
        );
    `); err != nil {
		log.Fatal("creating blocks table:", err)
	}
}

// — Workspace CRUD —

type Workspace struct {
	ID          int
	Name        string
	Description string
}

func CreateWorkspace(name, desc string) error {
	_, err := DB.Exec("INSERT INTO workspaces(name,description) VALUES(?,?)", name, desc)
	return err
}

func ListWorkspaces() ([]Workspace, error) {
	rows, err := DB.Query("SELECT id,name,description FROM workspaces")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var out []Workspace
	for rows.Next() {
		var w Workspace
		if err := rows.Scan(&w.ID, &w.Name, &w.Description); err != nil {
			return nil, err
		}
		out = append(out, w)
	}
	return out, nil
}

// — Stack CRUD —

type Stack struct {
	ID          int
	Name        string
	WorkspaceID int
}

func CreateStack(name string, workspaceID int) error {
	_, err := DB.Exec("INSERT INTO stacks(name,workspace_id) VALUES(?,?)", name, workspaceID)
	return err
}

func ListStacks(workspaceID int) ([]Stack, error) {
	rows, err := DB.Query("SELECT id,name,workspace_id FROM stacks WHERE workspace_id=?", workspaceID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var out []Stack
	for rows.Next() {
		var s Stack
		if err := rows.Scan(&s.ID, &s.Name, &s.WorkspaceID); err != nil {
			return nil, err
		}
		out = append(out, s)
	}
	return out, nil
}

// — Element CRUD —

type Element struct {
	ID          int
	Name        string
	Description string
}

func CreateElement(name, desc string) error {
	_, err := DB.Exec("INSERT INTO elements(name,description) VALUES(?,?)", name, desc)
	return err
}

func ListElements() ([]Element, error) {
	rows, err := DB.Query("SELECT id,name,description FROM elements")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var out []Element
	for rows.Next() {
		var e Element
		if err := rows.Scan(&e.ID, &e.Name, &e.Description); err != nil {
			return nil, err
		}
		out = append(out, e)
	}
	return out, nil
}

// — Block CRUD —

type Block struct {
	ID      int
	Name    string
	Content string
	StackID int
}

func CreateBlock(name, content string, stackID int) error {
	_, err := DB.Exec("INSERT INTO blocks(name,content,stack_id) VALUES(?,?,?)", name, content, stackID)
	return err
}

func ListBlocks(stackID int) ([]Block, error) {
	rows, err := DB.Query("SELECT id,name,content,stack_id FROM blocks WHERE stack_id=?", stackID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var out []Block
	for rows.Next() {
		var b Block
		if err := rows.Scan(&b.ID, &b.Name, &b.Content, &b.StackID); err != nil {
			return nil, err
		}
		out = append(out, b)
	}
	return out, nil
}
