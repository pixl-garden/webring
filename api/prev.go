package handler

import (
	"net/http"
	
	"github.com/pixl-garden/webring/pkg/database"
)

func PrevHandler(w http.ResponseWriter, r *http.Request) {
	db := database.GetDBClient()
	if db == nil {
		http.Error(w, "Database client is nil", http.StatusInternalServerError)
		return
	}

	currentSite := r.URL.Query().Get("site")
	prevSite, err := database.GetAdjacentSite(db, currentSite, false)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, prevSite, http.StatusFound)
}

