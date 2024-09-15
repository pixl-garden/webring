package handler

import (
	"context"
	"net/http"
	"sort"

	"github.com/pixl-garden/webring/internal/models"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	db := getDBClient()

	currentSite := r.URL.Query().Get("site")
	nextSite, err := getAdjacentSite(db, currentSite, true)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, nextSite, http.StatusFound)
}

func getAdjacentSite(db *db.Client, currentSite string, next bool) (string, error) {
	ctx := context.Background()
	ref := db.NewRef("members")
	var members map[string]models.Member
	if err := ref.Get(ctx, &members); err != nil {
		return "", err
	}

	var sites []string
	for _, member := range members {
		sites = append(sites, member.Website)
	}
	sort.Strings(sites)

	for i, site := range sites {
		if site == currentSite {
			if next {
				return sites[(i+1)%len(sites)], nil
			}
			return sites[(i-1+len(sites))%len(sites)], nil
		}
	}

	return sites[0], nil
}