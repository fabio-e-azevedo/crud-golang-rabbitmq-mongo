package mongodb

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	ph "crud-golang-rabbitmq-mongo/jsonplaceholder"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func FindOne(name string, value int) (*ph.User, error) {
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

	var result ph.User

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

func FindAll() (*[]ph.User, error) {
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

	var data []ph.User
	if err := cur.All(context.Background(), &data); err != nil {
		return nil, err
	}

	return &data, nil
}
