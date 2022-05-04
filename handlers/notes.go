package handlers

import (
	"database/sql"
	"fmt"
	"go-notes/data"
	"log"
	"net/http"
)

type Env struct {
	DB *sql.DB
}

// Define GetAll as a method on Env.
func (env *Env) GetAll(w http.ResponseWriter, r *http.Request) {
	// We can now access the connection pool directly in our handlers.
	nts, err := data.AllNotes(env.DB)
	if err != nil {
		log.Println(err)
		http.Error(w, http.StatusText(500), 500)
		return
	}

	for _, nt := range nts {
		fmt.Fprintf(w, "%d, %s\n", nt.ID, nt.Contents)
	}
}
