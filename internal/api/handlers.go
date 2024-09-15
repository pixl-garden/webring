package api

import (
	"context"
	"net/http"
	"time"

	"firebase.google.com/go/db"
	"github.com/pixl-garden/webring/internal/models"
)

type Handler struct {
	db *db.Client
}

func NewHandler(db *db.Client) *Handler {
	return &Handler{db: db}
}

func (h *Handler) Home(w http.ResponseWriter, r *http.Request) {
	// Implementation
}

func (h *Handler) Next(w http.ResponseWriter, r *http.Request) {
	// Implementation
}

func (h *Handler) Prev(w http.ResponseWriter, r *http.Request) {
	// Implementation
}

func (h *Handler) Members(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.getMembers(w, r)
	case http.MethodPost:
		h.addMember(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (h *Handler) getMembers(w http.ResponseWriter, r *http.Request) {
	// Implementation
}

func (h *Handler) addMember(w http.ResponseWriter, r *http.Request) {
	var member models.Member
	// Parse JSON from request body into member struct
	member.DateJoined = time.Now()

	ctx := context.Background()
	ref := h.db.NewRef("members")
	_, err := ref.Push(ctx, member)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
