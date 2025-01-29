package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Hello from Snippetbox"))
}

func main() {
    // serverMux is a go router
    mux := http.NewServeMux()
    // home is now the handler for the '/' route
    mux.HandleFunc("/", home)

    log.Print("Starting server on :4000")

    err := http.ListenAndServe(":4000", mux);
    log.Fatal(err);
}
