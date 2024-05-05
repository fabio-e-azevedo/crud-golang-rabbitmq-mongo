package mongodb

import (
	"context"
	"sync"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/*
Used to create a singleton object of MongoDB client.

Initialized and exposed through GetMongoClient().
*/
var singletonClient *mongo.Client

// Used during creation of singleton client object in GetMongoClient().
var clientInstanceError error

// Used to execute client creation procedure only once.
var mongoOnce sync.Once

type DbConnect struct {
	URI        string
	Database   string
	Collection string
	Client     *mongo.Client
}

func GetMongoClient(cfg *DbConnect) error {
	//Perform connection creation operation only once.
	mongoOnce.Do(func() {

		clientOptions := options.Client().ApplyURI(cfg.URI)

		client, err := mongo.Connect(context.Background(), clientOptions)
		if err != nil {
			clientInstanceError = err
		}

		err = client.Ping(context.Background(), nil)
		if err != nil {
			clientInstanceError = err
		}
		singletonClient = client
	})
	cfg.Client = singletonClient
	return clientInstanceError
}
