package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// Handler for the root route
	http.HandleFunc("/", helloWorldHandler)
	
	// Handler for a health check route
	http.HandleFunc("/health", healthHandler)
	
	// Configure the server port
	port := ":3000"
	fmt.Printf("Server running on port %s\n", port)
	fmt.Printf("Access: http://localhost%s\n", port)

	// Start the server
	log.Fatal(http.ListenAndServe(port, nil))
}

// Handler for Hello World
func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World! 🌍\n")
	fmt.Fprintf(w, "Method: %s\n", r.Method)
	fmt.Fprintf(w, "URL: %s\n", r.URL.Path)
}

// Handler for health check
func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, `{"status": "ok", "message": "Server is running!"}`)
}