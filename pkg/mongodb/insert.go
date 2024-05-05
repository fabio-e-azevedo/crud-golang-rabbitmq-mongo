package mongodb

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (m DbConnect) DbInsert(document interface{}) (string, error) {
	client, err := getMongoClient(&m)
	if err != nil {
		return "", err
	}

	coll := client.Database(m.Database).Collection(m.Collection)

	result, err := coll.InsertOne(context.TODO(), document)
	if err != nil {
		return "", err
	}

	return result.InsertedID.(primitive.ObjectID).Hex(), nil
}
