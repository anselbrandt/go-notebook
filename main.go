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

	gohandlers "github.com/gorilla/handlers"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/gorilla/mux"

	_ "github.com/mattn/go-sqlite3"
)

var bindAddress = env.String("BIND_ADDRESS", false, ":8080", "Bind address for the server")

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

	m.Use(middleware.Logger)

	get := m.Methods(http.MethodGet).Subrouter()

	get.HandleFunc("/notes", notes.GetAll)

	post := m.Methods(http.MethodPost).Subrouter()
	post.HandleFunc("/notes", notes.Add)

	put := m.Methods(http.MethodPut).Subrouter()
	put.HandleFunc("/notes", notes.Update)

	delete := m.Methods(http.MethodDelete).Subrouter()
	delete.HandleFunc("/notes/{id:[0-9]+}", notes.Delete)

	get.HandleFunc("/api/health", handlers.HealthCheck)

	spa := &handlers.SpaHandler{StaticPath: "frontend/build", IndexPath: "index.html"}

	m.PathPrefix("/").Handler(spa)

	// CORS
	// ch := gohandlers.CORS(gohandlers.AllowedOrigins([]string{"http://localhost:3000", "https://anselbrandt.dev", "https://www.anselbrandt.dev"}))
	methods := gohandlers.AllowedMethods([]string{"OPTIONS", "DELETE", "GET", "HEAD", "POST", "PUT", "PATCH"})
	origins := gohandlers.AllowedOrigins([]string{"*"})
	headers := gohandlers.AllowedHeaders([]string{"Content-Type"})
	ch := gohandlers.CORS(methods, origins, headers)

	// create a new server
	s := http.Server{
		Addr:         *bindAddress,      // configure the bind address
		Handler:      ch(m),             // set the default handler
		ErrorLog:     l,                 // set the logger for the server
		ReadTimeout:  5 * time.Second,   // max time to read request from the client
		WriteTimeout: 10 * time.Second,  // max time to write response to the client
		IdleTimeout:  120 * time.Second, // max time for connections using TCP Keep-Alive
	}

	// start the server
	go func() {
		l.Println("Starting server on port 8080")

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
