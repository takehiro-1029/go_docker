package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	_ "github.com/go-sql-driver/mysql"
)

func run() error {

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})
	r.Get("/welcome", func(w http.ResponseWriter, r *http.Request) {
		test(w, r)
	})
	http.ListenAndServe(":8080", r)

	return fmt.Errorf("dd")
}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func test(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("何もないおおおおおお"))
}
