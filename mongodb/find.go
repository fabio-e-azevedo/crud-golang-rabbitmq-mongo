package mongodb

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	jph "crud-golang-rabbitmq-mongo/jsonplaceholder"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func FindOne[T jph.ResourceGeneric](resource *T, name string, value int, collection string) (*T, error) {
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
	coll := client.Database(mongoDB).Collection(collection)

	filter := bson.M{name: value}

	err = coll.FindOne(context.TODO(), filter).Decode(resource)
	if err == mongo.ErrNoDocuments {
		fmt.Printf("No document was found with the title %s\n", filter)
		return nil, err
	}

	if err != nil {
		return nil, err
	}

	return resource, nil
}

func FindAll[T jph.ResourceGeneric](resource *[]T, collection string) (*[]T, error) {
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
	db := client.Database(mongoDB)

	filter := bson.M{}
	cur, err := db.Collection(collection).Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}

	if err := cur.All(context.Background(), resource); err != nil {
		return nil, err
	}

	return resource, nil
}
