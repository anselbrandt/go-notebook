package data

import (
	"database/sql"
)

type Note struct {
	ID       int
	Contents string
}

func Init(db *sql.DB) error {
	stmt, err := db.Prepare(`
	CREATE TABLE IF NOT EXISTS "notes" (
		"id" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		"contents" TEXT
	);
	`)
	if err != nil {
		return err
	}
	stmt.Exec()
	return nil
}

func AllNotes(db *sql.DB) ([]Note, error) {
	rows, err := db.Query("SELECT * FROM notes")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var nts []Note

	for rows.Next() {
		var nt Note

		err := rows.Scan(&nt.ID, &nt.Contents)
		if err != nil {
			return nil, err
		}

		nts = append(nts, nt)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return nts, nil
}
