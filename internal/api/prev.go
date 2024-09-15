package handler

import (
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	db := getDBClient()

	currentSite := r.URL.Query().Get("site")
	prevSite, err := getAdjacentSite(db, currentSite, false)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, prevSite, http.StatusFound)
}
