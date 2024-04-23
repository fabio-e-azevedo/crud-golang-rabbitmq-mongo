package mongodb

import (
	"context"
	"log"
	"os"

	jph "crud-golang-rabbitmq-mongo/jsonplaceholder"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Insert(body []byte, resourceType string) {
	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		log.Fatal("You must set your 'MONGODB_URI' environment variable.")
	}

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	switch resourceType {
	case "users":
		resource := (&jph.User{}).New(body)
		dataGeneric(resource, *client, resourceType)
	case "photos":
		resource := (&jph.Photo{}).New(body)
		dataGeneric(resource, *client, resourceType)
	case "posts":
		resource := (&jph.Posts{}).New(body)
		dataGeneric(resource, *client, resourceType)
	}
}

func dataGeneric[T jph.ResourceGeneric](resource T, client mongo.Client, resourceType string) {
	mongoDB := os.Getenv("MONGODB_DATABASE")
	coll := client.Database(mongoDB).Collection(resourceType)

	_, err := coll.InsertOne(context.TODO(), resource)
	if err != nil {
		panic(err)
	}

	//fmt.Println(result)
}
