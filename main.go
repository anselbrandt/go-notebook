package main

import (
	"database/sql"
	"log"
	"net/http"

	"go-notes/handlers"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "./notes.db")
	if err != nil {
		log.Fatal(err)
	}

	// Create an instance of Env containing the connection pool.
	env := &handlers.Env{DB: db}

	// Use env.GetAll as the handler function for the /books route.
	http.HandleFunc("/notes", env.GetAll)
	http.ListenAndServe(":3000", nil)
}
