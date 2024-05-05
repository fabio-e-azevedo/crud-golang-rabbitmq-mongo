package mongodb

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func FindOne[T any](resource *T, name string, value int, cfg *DbConnect) error {
	log.SetPrefix("[MNG] ")

	err := GetMongoClient(cfg)
	if err != nil {
		return err
	}

	coll := cfg.Client.Database(cfg.Database).Collection(cfg.Collection)

	filter := bson.M{name: value}

	err = coll.FindOne(context.Background(), filter).Decode(*resource)
	if err == mongo.ErrNoDocuments {
		log.Printf("no document was found with the field: \"%s\" and value: %d\n", name, value)
		return err
	}

	if err != nil {
		return err
	}

	return nil
}

func FindAll[T any](resource *[]T, cfg *DbConnect) error {
	log.SetPrefix("[MNG] ")

	// ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	// defer cancel()
	// client, err := mongo.Connect(ctx, options.Client().ApplyURI(m.URI))

	err := GetMongoClient(cfg)
	if err != nil {
		return err
	}

	coll := cfg.Client.Database(cfg.Database).Collection(cfg.Collection)

	filter := bson.M{}

	cur, err := coll.Find(context.Background(), filter)
	if err != nil {
		return err
	}

	if err := cur.All(context.Background(), resource); err != nil {
		return err
	}

	return nil
}

func FindAndDelete(id int, cfg *DbConnect) error {
	log.SetPrefix("[MNG] ")

	err := GetMongoClient(cfg)
	if err != nil {
		return err
	}

	coll := cfg.Client.Database(cfg.Database).Collection(cfg.Collection)

	filter := bson.M{"id": id}

	err = coll.FindOneAndDelete(context.Background(), filter).Err()
	if err != nil {
		return err
	}

	return nil
}
