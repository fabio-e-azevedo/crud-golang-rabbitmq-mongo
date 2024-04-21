package mongodb

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"crud-golang-rabbitmq-mongo/internal"
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

	var data jph.Resource[jph.User]
	json.Unmarshal(body, &data)

	switch resourceType {
	case "users":
		dataGeneric[jph.User](body, *client, resourceType)
	case "photos":
		dataGeneric[jph.Photo](body, *client, resourceType)
	}
}

func dataGeneric[T jph.ResourceGeneric](resourceBody []byte, client mongo.Client, resourceType string) {
	var resource T
	err := json.Unmarshal(resourceBody, &resource)
	internal.FailOnError(err, "Error Unmarshal resource")

	mongoDB := os.Getenv("MONGODB_DATABASE")
	coll := client.Database(mongoDB).Collection(resourceType)

	result, err := coll.InsertOne(context.TODO(), resource)
	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}
