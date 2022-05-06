package data

import (
	"database/sql"
	"time"
)

type Note struct {
	ID        int
	Contents  string
	CreatedAt int
	UpdatedAt int
}

type NoteStore struct {
	DB *sql.DB
}

func (store NoteStore) Init() error {
	stmt, err := store.DB.Prepare(`
	CREATE TABLE IF NOT EXISTS "notes" (
		"id" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		"contents" TEXT,
		"createdAt" INTEGER,
		"updatedAt" INTEGER
	);
	`)
	if err != nil {
		return err
	}
	stmt.Exec()
	return nil
}

func (store NoteStore) GetAll() ([]Note, error) {
	rows, err := store.DB.Query("SELECT * FROM notes")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var nts []Note

	for rows.Next() {
		var nt Note

		err := rows.Scan(&nt.ID, &nt.Contents, &nt.CreatedAt, &nt.UpdatedAt)
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

func (store NoteStore) Get(rowid int64) (Note, error) {
	note := Note{}
	var id int
	var contents string
	var createdAt int
	var updatedAt int
	err := store.DB.QueryRow(`
	SELECT * FROM notes
	WHERE id=?`, rowid).Scan(&id, &contents, &createdAt, &updatedAt)
	if err != nil {
		return note, err
	}
	note = Note{
		ID:        id,
		Contents:  contents,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}
	return note, nil
}

func (store NoteStore) Add(note Note) (int64, error) {
	stmt, err := store.DB.Prepare(`
	INSERT INTO notes (contents, createdAt, UpdatedAt) values (?, ?, ?)
	`)
	if err != nil {
		return 0, err
	}
	t := time.Now().Unix()
	result, err := stmt.Exec(note.Contents, t, t)
	if err != nil {
		return 0, err
	}
	rowid, err := result.LastInsertId()
	if err != nil {
		return rowid, err
	}
	return rowid, nil
}

func (store NoteStore) Update(note Note) (int64, error) {
	stmt, err := store.DB.Prepare(`
	UPDATE notes SET contents=?, updatedAt=? WHERE id=?
	`)
	if err != nil {
		return 0, err
	}
	t := time.Now().Unix()
	result, err := stmt.Exec(note.Contents, t, note.ID)
	if err != nil {
		return 0, err
	}
	rowid, err := result.LastInsertId()
	if err != nil {
		return rowid, err
	}
	return rowid, nil
}

func (store NoteStore) Touch(note Note) (int64, error) {
	stmt, err := store.DB.Prepare(`
	UPDATE notes SET updatedAt=? WHERE id=?
	`)
	if err != nil {
		return 0, err
	}
	t := time.Now().Unix()
	result, err := stmt.Exec(t, note.ID)
	if err != nil {
		return 0, err
	}
	rowid, err := result.LastInsertId()
	if err != nil {
		return rowid, err
	}
	return rowid, nil
}

func (store NoteStore) Delete(id int) (int64, error) {
	stmt, err := store.DB.Prepare(`
	DELETE FROM notes WHERE id=?
	`)
	if err != nil {
		return 0, err
	}
	result, err := stmt.Exec(id)
	if err != nil {
		return 0, err
	}
	rowid, err := result.LastInsertId()
	if err != nil {
		return rowid, err
	}
	return rowid, nil
}
