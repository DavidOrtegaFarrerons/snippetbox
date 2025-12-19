package main

import (
	"log"
	"net/http"
	"snippetbox.davidortegafarrerons.com/internal/web"
)

func main() {
	mux := http.NewServeMux()

	fileServer := http.FileServer(
		web.NeuteredFileSystem(http.Dir("./ui/static")),
	)

	mux.Handle("/static", http.NotFoundHandler())
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)

	log.Print("Starting server on :4000")
	err := http.ListenAndServe("localhost:4000", mux)
	log.Fatal(err)
}
