package handler

import (
	"net/http"

	"github.com/pixl-garden/webring/pkg/database"
)

func NextHandler(w http.ResponseWriter, r *http.Request) {
	db := database.GetDBClient()

	currentSite := r.URL.Query().Get("site")
	nextSite, err := database.GetAdjacentSite(db, currentSite, true)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, nextSite, http.StatusFound)
}