package handlers

import (
	"encoding/json"
	"go-notes/data"
	"log"
	"net/http"
)

type Notebook struct {
	Notes data.NoteStore
}

// Define GetAll as a method on Notebook.
func (notebook *Notebook) GetAll(w http.ResponseWriter, r *http.Request) {
	// We can now access the connection pool directly in our handlers.
	notes, err := notebook.Notes.AllNotes()
	if err != nil {
		log.Println(err)
		http.Error(w, http.StatusText(500), 500)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(notes)
}
