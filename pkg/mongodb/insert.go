package mongodb

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (m DbConnect) DbInsert(document interface{}) (string, error) {
	err := GetMongoClient(&m)
	if err != nil {
		return "", err
	}

	coll := m.Client.Database(m.Database).Collection(m.Collection)

	result, err := coll.InsertOne(context.Background(), document)
	if err != nil {
		return "", err
	}

	return result.InsertedID.(primitive.ObjectID).Hex(), nil
}
