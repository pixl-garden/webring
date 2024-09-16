package handler

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/pixl-garden/webring/pkg/models"
	"github.com/pixl-garden/webring/pkg/database"
	"github.com/pixl-garden/webring/pkg/utils"
	"firebase.google.com/go/db" 
)

func MembersHandler(w http.ResponseWriter, r *http.Request) {
	db := database.GetDBClient()

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

	var allMembers map[string]models.Member
	if err := ref.Get(ctx, &allMembers); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	memberSlice := make([]models.Member, 0, len(allMembers))
	for _, m := range allMembers {
		memberSlice = append(memberSlice, m)
	}

	if err := utils.UpdateReadme(memberSlice); err != nil {
		log.Printf("Failed to update README: %v", err)
	}
}
