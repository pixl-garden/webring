package handler

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"runtime"

	"github.com/pixl-garden/webring/pkg/database"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	log.Println("Handler called")
	log.Printf("Go version: %s", runtime.Version())
	
	log.Println("FIREBASE_DATABASE_URL:", os.Getenv("FIREBASE_DATABASE_URL"))
	log.Println("FIREBASE_CREDENTIALS length:", len(os.Getenv("FIREBASE_CREDENTIALS")))

	db := database.GetDBClient()
	if db == nil {
		log.Println("Database client is nil")
	} else {
		log.Println("Database client initialized successfully")
	}

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	fmt.Fprintf(w, "pixl garden webring!") // TODO: add more info
}