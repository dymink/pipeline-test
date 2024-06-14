package main

import (
	// "database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq"
)

func main() {
	// Connect to the database
	// dbURL := os.Getenv("DATABASE_URL")
	// db, err := sql.Open("postgres", dbURL)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer db.Close()

	// Serve static files
	fs := http.FileServer(http.Dir("./frontend/dist"))
	http.Handle("/", fs)

	// Example endpoint
	http.HandleFunc("/api/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello from Go!")
	})

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Listening on port %s...", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
