package database

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
)

func Connect() (*mongo.Database, error) {
	client, err := mongo.NewClient(options.Client().ApplyURI(os.Getenv("DATABASE_HOST")))
	if err != nil {
		return nil, err
	}

	err = client.Connect(context.TODO())
	if err != nil {
		return nil, err
	}

	return client.Database(os.Getenv("DATABASE_NAME")), nil
}
