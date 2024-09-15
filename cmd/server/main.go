package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/pixl-garden/webring/internal/api"
	"github.com/pixl-garden/webring/internal/database"
)

func main() {
	db, err := database.InitFirebase()
	if err != nil {
		log.Fatalf("Error initializing database: %v", err)
	}

	handler := api.NewHandler(db)

	http.HandleFunc("/", handler.Home)
	http.HandleFunc("/next", handler.Next)
	http.HandleFunc("/prev", handler.Prev)
	http.HandleFunc("/api/members", handler.Members)

	fmt.Println("Server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}