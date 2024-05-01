package mongodb

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DbConnect struct {
	URI        string
	Database   string
	Collection string
}

func (m DbConnect) DbInsert(document interface{}) (string, error) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(m.URI))
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	coll := client.Database(m.Database).Collection(m.Collection)

	result, err := coll.InsertOne(context.TODO(), document)
	if err != nil {
		panic(err)
	}

	return result.InsertedID.(primitive.ObjectID).Hex(), nil
}
