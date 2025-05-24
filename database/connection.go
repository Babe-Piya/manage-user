package database

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.mongodb.org/mongo-driver/v2/mongo/readpref"
)

func NewConnection(user, password, host, port, databaseName string) (*mongo.Database, error) {
	url := fmt.Sprintf("mongodb://%s:%s@%s:%s", user, password, host, port)
	client, err := mongo.Connect(options.Client().ApplyURI(url))
	if err != nil {
		return nil, err
	}

	err = client.Ping(context.Background(), readpref.Primary())
	if err != nil {
		return nil, err
	}

	db := client.Database(databaseName)

	return db, nil
}
