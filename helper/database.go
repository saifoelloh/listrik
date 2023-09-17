package helper

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func DatabaseConnection() (*mongo.Database, context.CancelFunc, error) {
	user := GoDotEnvVariable("DB_USERNAME")
	pass := GoDotEnvVariable("DB_PASSWORD")
	host := GoDotEnvVariable("DB_HOST")
	port := GoDotEnvVariable("DB_PORT")
	databaseName := GoDotEnvVariable("DB_NAME")
	mongoURI := "mongodb://" + user + ":" + pass + "@" + host + ":" + port + "/?authSource=admin"

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	mongoOption := options.Client().ApplyURI(mongoURI).SetServerSelectionTimeout(5 * time.Second)
	client, err := mongo.Connect(ctx, mongoOption)
	if err != nil {
		cancel()
		return nil, nil, err
	}

	db := client.Database(databaseName)
	return db, cancel, nil
}
