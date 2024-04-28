package mongodb

import (
	"context"
	"fmt"

	jph "crud-golang-rabbitmq-mongo/jsonplaceholder"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DbConnect struct {
	URI        string
	Database   string
	Collection string
}

func (m DbConnect) DbInsert(body []byte) string {
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

	resource, err := jph.GetResource(body)
	if err != nil {
		return fmt.Sprintln(err)
	}

	result, err := coll.InsertOne(context.TODO(), resource)
	if err != nil {
		panic(err)
	}

	return result.InsertedID.(primitive.ObjectID).Hex()
}
