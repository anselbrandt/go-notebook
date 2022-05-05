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
	notes, err := notebook.Notes.GetAll()
	if err != nil {
		log.Println(err)
		http.Error(w, http.StatusText(500), 500)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(notes)
}

func (notebook *Notebook) Add(w http.ResponseWriter, r *http.Request) {
	var n data.Note
	err := json.NewDecoder(r.Body).Decode(&n)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "", http.StatusBadRequest)
		return
	}
	if len(n.Contents) == 0 {
		http.Error(w, "", http.StatusBadRequest)
		return
	}
	contents := data.Note{
		Contents: n.Contents,
	}
	rowid, addErr := notebook.Notes.Add(contents)
	if addErr != nil {
		log.Println(addErr.Error())
		http.Error(w, "", http.StatusInternalServerError)
		return
	}
	note, getErr := notebook.Notes.Get(rowid)
	if getErr != nil {
		log.Println(getErr.Error())
		http.Error(w, "", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(note)

}

func (notebook *Notebook) Update(w http.ResponseWriter, r *http.Request) {
	var n data.Note
	err := json.NewDecoder(r.Body).Decode(&n)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	var rowid int64
	var addErr error

	if len(n.Contents) == 0 {
		rowid, addErr = notebook.Notes.Touch(n)
	} else {
		rowid, addErr = notebook.Notes.Update(n)
	}

	if addErr != nil {
		log.Println(addErr.Error())
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	note, getErr := notebook.Notes.Get(rowid)
	if getErr != nil {
		log.Println(getErr.Error())
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(note)

}
