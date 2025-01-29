package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func home(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Hello from Snippetbox"))
}

func snippetView(w http.ResponseWriter, r *http.Request) {
    // convert to integer, returns an error if not possible
    id, err := strconv.Atoi(r.PathValue("id"))
    if (err != nil || id < 1) {
        http.NotFound(w, r)
        return
    }

    msg := fmt.Sprintf("Display a specific snippet with ID %d...", id)

    w.Write([]byte(msg))
}

func snippetCreate(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Display a form for creating a new snippet..."))
}

func main() {
    // serverMux is a go router
    mux := http.NewServeMux()
    // home is now the handler for the '/' route
    mux.HandleFunc("GET /{$}", home) // The $ stops it from being a catch-all
    mux.HandleFunc("GET /snippet/view/{id}", snippetView)
    mux.HandleFunc("GET /snippet/create", snippetCreate)

    log.Print("Starting server on :4000")

    err := http.ListenAndServe(":4000", mux);
    log.Fatal(err);
}
