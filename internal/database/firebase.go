package database

import (
	"context"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/db"
	"google.golang.org/api/option"
)

func InitFirebase() (*db.Client, error) {
	ctx := context.Background()
	conf := &firebase.Config{
		DatabaseURL: "https://your-project-id.firebaseio.com",
	}
	opt := option.WithCredentialsFile("path/to/your/service-account-key.json")
	app, err := firebase.NewApp(ctx, conf, opt)
	if err != nil {
		return nil, err
	}

	return app.Database(ctx)
}
