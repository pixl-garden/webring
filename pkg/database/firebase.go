package database

import (
	"context"
	"log"
	"os"
	"sync"
	"sort"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/db"
	"google.golang.org/api/option"

	"github.com/pixl-garden/webring/pkg/models"
)

var (
	dbClient *db.Client
	once     sync.Once
)

func initFirebase() {
	once.Do(func() {
		ctx := context.Background()
		conf := &firebase.Config{
			DatabaseURL: os.Getenv("FIREBASE_DATABASE_URL"),
		}
		
		// Use Firebase credentials from environment variable
		opt := option.WithCredentialsJSON([]byte(os.Getenv("FIREBASE_CREDENTIALS")))
		
		app, err := firebase.NewApp(ctx, conf, opt)
		if err != nil {
			log.Fatalf("Error initializing app: %v", err)
		}

		dbClient, err = app.Database(ctx)
		if err != nil {
			log.Fatalf("Error initializing database client: %v", err)
		}
	})
}

func GetDBClient() *db.Client {
	initFirebase()
	return dbClient
}

func GetAdjacentSite(db *db.Client, currentSite string, next bool) (string, error) {
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