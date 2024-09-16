package handler

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/pixl-garden/webring/internal/models"
	"github.com/pixl-garden/webring/internal/utils"
	"github.com/pixl-garden/webring/internal/database"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	db := getDBClient()

	switch r.Method {
	case http.MethodGet:
		getMembers(w, r, db)
	case http.MethodPost:
		addMember(w, r, db)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func getMembers(w http.ResponseWriter, r *http.Request, db *db.Client) {
	ctx := context.Background()
	ref := db.NewRef("members")
	var members map[string]models.Member
	if err := ref.Get(ctx, &members); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(members)
}

func addMember(w http.ResponseWriter, r *http.Request, db *db.Client) {
	var member models.Member
	if err := json.NewDecoder(r.Body).Decode(&member); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	member.DateJoined = time.Now()

	ctx := context.Background()
	ref := db.NewRef("members")
	_, err := ref.Push(ctx, member)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(member)

	var allMembers []models.Member
	if err := ref.Get(ctx, &allMembers); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := utils.UpdateReadme(allMembers); err != nil {
		// Log the error, but don't return it to the client
		log.Printf("Failed to update README: %v", err)
	}
}
