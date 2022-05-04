package data

import (
	"database/sql"
)

type Note struct {
	ID       int
	Contents string
}

type NoteStore struct {
	DB *sql.DB
}

func (store NoteStore) Init() error {
	stmt, err := store.DB.Prepare(`
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

func (store NoteStore) AllNotes() ([]Note, error) {
	rows, err := store.DB.Query("SELECT * FROM notes")
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
