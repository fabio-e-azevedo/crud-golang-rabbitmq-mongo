package mongodb

import (
	"context"
	"encoding/json"
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

	log.SetPrefix("")
	return nil
}

func FindAll(cfg *DbConnect) ([]byte, int, error) {
	log.SetPrefix("[MNG] ")

	var totalDocuments int

	err := GetMongoClient(cfg)
	if err != nil {
		return nil, totalDocuments, err
	}

	coll := cfg.Client.Database(cfg.Database).Collection(cfg.Collection)

	cur, err := coll.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, totalDocuments, err
	}

	var documents []bson.M
	if err := cur.All(context.Background(), &documents); err != nil {
		return nil, totalDocuments, err
	}

	documentsBytes, err := json.Marshal(documents)
	if err != nil {
		return nil, totalDocuments, err
	}

	totalDocuments = len(documents)

	log.SetPrefix("")
	return documentsBytes, totalDocuments, nil
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

	log.SetPrefix("")
	return nil
}
