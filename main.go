package main

import (
	"github.com/kutaui/htmx-go/handler"
	"log"
	"net/http"
)

func main() {
	homeHandler := handler.HomeHandler{}
	http.HandleFunc("/", homeHandler.ServeHTTP)

	log.Println("Server is starting on :8080...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
