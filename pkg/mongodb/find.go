package mongodb

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func FindOne[T any](resource *T, name string, value int, m *DbConnect) error {
	log.SetPrefix("[MNG] ")

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(m.URI))
	if err != nil {
		return err
	}

	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	coll := client.Database(m.Database).Collection(m.Collection)

	filter := bson.M{name: value}

	err = coll.FindOne(context.TODO(), filter).Decode(*resource)
	if err == mongo.ErrNoDocuments {
		log.Printf("no document was found with the field: \"%s\" and value: %d\n", name, value)
		return err
	}

	if err != nil {
		return err
	}

	return nil
}

func FindAll[T any](resource *[]T, m *DbConnect) error {
	log.SetPrefix("[MNG] ")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(m.URI))
	if err != nil {
		return err
	}

	db := client.Database(m.Database)

	filter := bson.M{}
	cur, err := db.Collection(m.Collection).Find(context.Background(), filter)
	if err != nil {
		return err
	}

	if err := cur.All(context.Background(), resource); err != nil {
		return err
	}

	return nil
}

func FindAndDelete(id int, m *DbConnect) error {
	log.SetPrefix("[MNG] ")

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(m.URI))
	if err != nil {
		return err
	}

	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	filter := bson.M{"id": id}
	coll := client.Database(m.Database).Collection(m.Collection)

	err = coll.FindOneAndDelete(context.TODO(), filter).Err()
	if err != nil {
		return err
	}

	return nil
}
