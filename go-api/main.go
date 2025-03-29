package main

import (
	"fmt"
	"net/http"
	"github.com/rs/cors"
)

// Handler for the root endpoint
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello from Go API!")
}

func main() {
	// Set up CORS handler
	corsHandler := cors.New(cors.Options{
		AllowedOrigins: []string{"*"}, // Allow all origins, adjust this for more secure access
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Content-Type"},
	})

	// Register the handler
	http.HandleFunc("/", handler)

	// Wrap the default mux with CORS handler
	handlerWithCORS := corsHandler.Handler(http.DefaultServeMux)

	// Start the server on port 8080
	http.ListenAndServe(":8080", handlerWithCORS)
}
