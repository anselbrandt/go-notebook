package main

import (
	"context"
	"database/sql"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"go-notes/data"
	"go-notes/env"
	"go-notes/handlers"

	"github.com/gorilla/mux"

	_ "github.com/mattn/go-sqlite3"
)

var bindAddress = env.String("BIND_ADDRESS", false, ":9090", "Bind address for the server")

func main() {

	env.Parse()

	l := log.New(os.Stdout, "go-notes ", log.LstdFlags)

	db, err := sql.Open("sqlite3", "./notes.db")
	if err != nil {
		log.Fatal(err)
	}

	store := data.NoteStore{DB: db}

	store.Init()

	// Create an instance of Notes containing the connection pool.
	notes := &handlers.Notebook{Notes: store}

	// Create new serve mux
	m := mux.NewRouter()

	get := m.Methods(http.MethodGet).Subrouter()

	get.HandleFunc("/notes", notes.GetAll)

	post := m.Methods(http.MethodPost).Subrouter()
	post.HandleFunc("/notes", notes.Add)

	get.HandleFunc("/api/health", handlers.HealthCheck)

	spa := &handlers.SpaHandler{StaticPath: "frontend/build", IndexPath: "index.html"}

	m.PathPrefix("/").Handler(spa)

	// create a new server
	s := http.Server{
		Addr:         *bindAddress,      // configure the bind address
		Handler:      m,                 // set the default handler
		ErrorLog:     l,                 // set the logger for the server
		ReadTimeout:  5 * time.Second,   // max time to read request from the client
		WriteTimeout: 10 * time.Second,  // max time to write response to the client
		IdleTimeout:  120 * time.Second, // max time for connections using TCP Keep-Alive
	}

	// start the server
	go func() {
		l.Println("Starting server on port 9090")

		err := s.ListenAndServe()
		if err != nil {
			l.Printf("Error starting server: %s\n", err)
			os.Exit(1)
		}
	}()

	// trap sigterm or interupt and gracefully shutdown the server
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, os.Kill)

	// Block until a signal is received.
	sig := <-c
	log.Println("Got signal:", sig)

	// gracefully shutdown the server, waiting max 30 seconds for current operations to complete
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(ctx)
}
