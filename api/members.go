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
	log.Println("MembersHandler called")
	db := database.GetDBClient()
	if db == nil {
		http.Error(w, "Database client is nil", http.StatusInternalServerError)
		log.Println("Database client is nil")
		return
	}

	switch r.Method {
	case http.MethodGet:
		getMembers(w, db)
	case http.MethodPost:
		addMember(w, r, db)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		log.Printf("Method not allowed: %v", r.Method)
	}
}

func getMembers(w http.ResponseWriter, db *db.Client) {
	log.Println("getMembers called")
	ctx := context.Background()
	ref := db.NewRef("members")
	var members map[string]models.Member
	if err := ref.Get(ctx, &members); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("Error getting members: %v", err)
		return
	}

	memberSlice := make([]models.Member, 0, len(members))
	for _, member := range members {
		memberSlice = append(memberSlice, member)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(memberSlice)
}

func addMember(w http.ResponseWriter, r *http.Request, db *db.Client) {
	log.Println("addMember called")
	var member models.Member
	if err := json.NewDecoder(r.Body).Decode(&member); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Printf("Error decoding member: %v", err)
		return
	}
	member.DateJoined = time.Now()

	ctx := context.Background()
	ref := db.NewRef("members")
	_, err := ref.Push(ctx, member)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("Error adding member: %v", err)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(member)

	var allMembers map[string]models.Member
	if err := ref.Get(ctx, &allMembers); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("Error getting members: %v", err)
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
