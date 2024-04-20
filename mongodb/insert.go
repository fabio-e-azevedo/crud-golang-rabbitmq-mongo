package mongodb

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	ph "crud-golang-rabbitmq-mongo/jsonplaceholder"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Insert(body []byte) {
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

	var data ph.Resource
	json.Unmarshal(body, &data)

	mongoDB := os.Getenv("MONGODB_DATABASE")
	coll := client.Database(mongoDB).Collection(data.ResourceType)

	result, err := coll.InsertOne(context.TODO(), data.Data)
	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}
