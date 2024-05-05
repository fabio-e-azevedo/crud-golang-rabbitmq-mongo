package mongodb

import (
	"context"
	"sync"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/*
Used to create a singleton object of MongoDB client.

Initialized and exposed through getMongoClient().
*/
var clientInstance *mongo.Client

// Used during creation of singleton client object in getMongoClient().
var clientInstanceError error

// Used to execute client creation procedure only once.
var mongoOnce sync.Once

type DbConnect struct {
	URI        string
	Database   string
	Collection string
}

// getMongoClient - Return mongodb connection to work with
func getMongoClient(m *DbConnect) (*mongo.Client, error) {
	//Perform connection creation operation only once.
	mongoOnce.Do(func() {

		clientOptions := options.Client().ApplyURI(m.URI)

		client, err := mongo.Connect(context.TODO(), clientOptions)
		if err != nil {
			clientInstanceError = err
		}

		err = client.Ping(context.TODO(), nil)
		if err != nil {
			clientInstanceError = err
		}
		clientInstance = client
	})
	return clientInstance, clientInstanceError
}
