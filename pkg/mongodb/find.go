package mongodb

import (
	"bytes"
	"context"
	"encoding/json"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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

func FindOneById(idNumber int, cfg *DbConnect) ([]byte, error) {
	log.SetPrefix("[MNG] ")

	err := GetMongoClient(cfg)
	if err != nil {
		return nil, err
	}

	coll := cfg.Client.Database(cfg.Database).Collection(cfg.Collection)

	filter := bson.M{"_id": idNumber}
	var documents bson.M

	err = coll.FindOne(context.Background(), filter).Decode(&documents)
	if err == mongo.ErrNoDocuments {
		return nil, err
	}

	if err != nil {
		return nil, err
	}

	documentsBytes, err := json.Marshal(documents)
	if err != nil {
		return nil, err
	}

	documentsBytes = bytes.Replace(documentsBytes, []byte("_id"), []byte("id"), -1)

	log.SetPrefix("")
	return documentsBytes, nil
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
	documentsBytes = bytes.Replace(documentsBytes, []byte("_id"), []byte("id"), -1)

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

	filter := bson.M{"_id": id}

	err = coll.FindOneAndDelete(context.Background(), filter).Err()
	if err != nil {
		return err
	}

	log.SetPrefix("")
	return nil
}

func LatestId(cfg *DbConnect) (int, error) {
	err := GetMongoClient(cfg)
	if err != nil {
		return 0, err
	}
	collection := cfg.Client.Database(cfg.Database).Collection(cfg.Collection)

	opts := options.FindOne().SetSort(map[string]int{"_id": -1})
	result := collection.FindOne(context.Background(), bson.M{}, opts)
	if err := result.Err(); err != nil {
		return 0, err
	}

	var document bson.M
	if err := result.Decode(&document); err != nil {
		return 0, err
	}
	lastID := document["_id"]

	resultInt32 := lastID.(int32)
	resultInt := int(resultInt32)

	return resultInt + 1, nil
}
