package handlers

import (
	"database/sql"
	"fmt"
	"go-notes/data"
	"log"
	"net/http"
)

type Notebook struct {
	DB *sql.DB
}

// Define GetAll as a method on Notebook.
func (notebook *Notebook) GetAll(w http.ResponseWriter, r *http.Request) {
	// We can now access the connection pool directly in our handlers.
	nts, err := data.AllNotes(notebook.DB)
	if err != nil {
		log.Println(err)
		http.Error(w, http.StatusText(500), 500)
		return
	}

	for _, nt := range nts {
		fmt.Fprintf(w, "%d, %s\n", nt.ID, nt.Contents)
	}
}
