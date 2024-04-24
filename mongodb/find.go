package mongodb

import (
	"context"
	"fmt"
	"time"

	jph "crud-golang-rabbitmq-mongo/jsonplaceholder"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func FindOne[T jph.ResourceGeneric](resource *T, name string, value int, m *DbConnect) (*T, error) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(m.URI))
	if err != nil {
		return nil, err
	}

	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	coll := client.Database(m.Database).Collection(m.Collection)

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

func FindAll[T jph.ResourceGeneric](resource *[]T, m *DbConnect) (*[]T, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(m.URI))
	if err != nil {
		return nil, err
	}

	db := client.Database(m.Database)

	filter := bson.M{}
	cur, err := db.Collection(m.Collection).Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}

	if err := cur.All(context.Background(), resource); err != nil {
		return nil, err
	}

	return resource, nil
}
