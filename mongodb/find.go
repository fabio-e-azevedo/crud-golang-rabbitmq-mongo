package mongodb

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"crud-golang-rabbitmq-mongo/internal"
	"crud-golang-rabbitmq-mongo/users"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func FindOne(name string, value int) (*users.UserFromBson, error) {
	err := godotenv.Load()
	internal.FailOnError(err, "No .env file found")

	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		log.Fatal("You must set your 'MONGODB_URI' environment variable.")
	}

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}

	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	mongoDB := os.Getenv("MONGODB_DATABASE")
	mongoCollection := os.Getenv("MONGODB_COLLECTION")
	coll := client.Database(mongoDB).Collection(mongoCollection)

	filter := bson.M{name: value}

	var result users.UserFromBson

	err = coll.FindOne(context.TODO(), filter).Decode(&result)
	if err == mongo.ErrNoDocuments {
		fmt.Printf("No document was found with the title %s\n", filter)
		return nil, err
	}

	if err != nil {
		return nil, err
	}

	return &result, nil
}

func FindAll() (*[]users.UserFromBson, error) {
	err := godotenv.Load()
	internal.FailOnError(err, "No .env file found")

	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		log.Fatal("You must set your 'MONGODB_URI' environment variable.")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}

	mongoDB := os.Getenv("MONGODB_DATABASE")
	mongoCollection := os.Getenv("MONGODB_COLLECTION")

	db := client.Database(mongoDB)

	filter := bson.M{}
	cur, err := db.Collection(mongoCollection).Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}

	var data []users.UserFromBson
	if err := cur.All(context.Background(), &data); err != nil {
		return nil, err
	}

	return &data, nil
}
