package handler

import (
	"context"
	"log"
	"os"
	"sync"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/db"
	"google.golang.org/api/option"
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