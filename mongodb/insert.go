package mongodb

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	user "crud-golang-rabbitmq-mongo/users"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Insert(body []byte) {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

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

	mongoDB := os.Getenv("MONGODB_DATABASE")
	mongoCollection := os.Getenv("MONGODB_COLLECTION")
	coll := client.Database(mongoDB).Collection(mongoCollection)

	var data user.UserFromJson

	json.Unmarshal(body, &data)

	result, err := coll.InsertOne(context.TODO(), data)
	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}
