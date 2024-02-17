package main

import (
	"github.com/a-h/templ"
	"github.com/kutaui/htmx-go/view/layout"
	"log"
	"net/http"
)

func main() {

	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", templ.Handler(layout.Base()))

	http.Handle("/static/", http.StripPrefix("/static/", fs))

	log.Println("Server is starting on :8080...")
	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		panic(err)
	}
}
