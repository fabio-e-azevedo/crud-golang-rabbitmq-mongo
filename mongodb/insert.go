package mongodb

import (
	"context"

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

	switch m.Collection {
	case "users":
		resource := (&jph.User{}).New(body)
		r := dataGeneric(resource, *coll)
		return r
	case "photos":
		resource := (&jph.Photo{}).New(body)
		r := dataGeneric(resource, *coll)
		return r
	case "posts":
		resource := (&jph.Posts{}).New(body)
		r := dataGeneric(resource, *coll)
		return r
	case "comments":
		resource := (&jph.Comments{}).New(body)
		r := dataGeneric(resource, *coll)
		return r
	}
	return ""
}

func dataGeneric[T jph.ResourceGeneric](resource T, coll mongo.Collection) string {
	result, err := coll.InsertOne(context.TODO(), resource)
	if err != nil {
		panic(err)
	}

	return result.InsertedID.(primitive.ObjectID).Hex()
}
