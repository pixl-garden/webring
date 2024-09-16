package handler

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"runtime"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	log.Println("Handler called")
	log.Printf("Go version: %s", runtime.Version())
	
	// Log environment variables
	log.Println("FIREBASE_DATABASE_URL:", os.Getenv("FIREBASE_DATABASE_URL"))
	log.Println("FIREBASE_CREDENTIALS length:", len(os.Getenv("FIREBASE_CREDENTIALS")))

	// Temporarily comment out database initialization for testing
	// db := database.GetDBClient()
	// if db == nil {
	// 	log.Println("Database client is nil")
	// } else {
	// 	log.Println("Database client initialized successfully")
	// }

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	fmt.Fprintf(w, "Hello from Pixl Garden Webring! Running on Go %s", runtime.Version())
}