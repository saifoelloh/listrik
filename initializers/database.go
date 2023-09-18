package initializers

import (
	"context"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client
var Cancel context.CancelFunc

func DatabaseConnection() error {
	var err error
	var ctx context.Context
	mongoURI := os.Getenv("DB_URI")

	ctx, Cancel = context.WithTimeout(context.Background(), 10*time.Second)
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	mongoOption := options.Client()
	mongoOption.
		ApplyURI(mongoURI).
		SetServerAPIOptions(serverAPI).
		SetServerSelectionTimeout(10 * time.Second)

	Client, err = mongo.Connect(ctx, mongoOption)
	if err != nil {
		Cancel()
		return err
	}

	return nil
}
