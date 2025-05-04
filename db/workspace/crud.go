package workspace

import (
	"fmt"
	"time"

	"nookli/db"
)

// Workspace mirrors pkg/service/workspace.Workspace but includes CreatedAt.
type Workspace struct {
	ID          int
	Name        string
	Description string
	CreatedAt   time.Time
}

func Create(name, desc string) error {
	_, err := db.DB.Exec("INSERT INTO workspaces(name,description,created_at) VALUES(?,?,?)",
		name, desc, time.Now())
	return err
}

func List() ([]Workspace, error) {
	rows, err := db.DB.Query("SELECT id,name,description,created_at FROM workspaces")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var out []Workspace
	for rows.Next() {
		var w Workspace
		if err := rows.Scan(&w.ID, &w.Name, &w.Description, &w.CreatedAt); err != nil {
			return nil, err
		}
		out = append(out, w)
	}
	return out, nil
}

func Get(id int) (*Workspace, error) {
	row := db.DB.QueryRow(
		"SELECT id,name,description,created_at FROM workspaces WHERE id = ?",
		id,
	)
	var w Workspace
	if err := row.Scan(&w.ID, &w.Name, &w.Description, &w.CreatedAt); err != nil {
		return nil, err
	}
	return &w, nil
}

func Update(id int, name, desc string) error {
	result, err := db.DB.Exec(
		"UPDATE workspaces SET name = ?, description = ? WHERE id = ?",
		name, desc, id,
	)
	if err != nil {
		return err
	}
	if n, _ := result.RowsAffected(); n == 0 {
		return fmt.Errorf("no rows updated")
	}
	return nil
}

func Delete(id int) error {
	result, err := db.DB.Exec("DELETE FROM workspaces WHERE id = ?", id)
	if err != nil {
		return err
	}
	if n, _ := result.RowsAffected(); n == 0 {
		return fmt.Errorf("no rows deleted")
	}
	return nil
}
