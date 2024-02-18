package main

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
	"github.com/kutaui/htmx-go/db/sqlc"
	"github.com/kutaui/htmx-go/handler"
	"log"
	"net/http"
	"os"
)

func main() {
	err := godotenv.Load()

	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	q := db.New(conn)

	author, err := q.GetCollection(context.Background(), 1)
	if err != nil {
		fmt.Fprintf(os.Stderr, "GetAuthor failed: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(author.Name)
	homeHandler := handler.HomeHandler{}
	registerPageHandler := handler.RegisterPageHandler{}
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.Handle("/", &homeHandler)
	http.Handle("/register", &registerPageHandler)
	log.Println("Server is starting on :8080...")
	http.ListenAndServe("localhost:8080", nil)
}
