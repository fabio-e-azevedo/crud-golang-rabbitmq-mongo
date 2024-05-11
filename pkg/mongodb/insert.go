package mongodb

import (
	"context"
	"fmt"
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

	return fmt.Sprintf("insert document id %v", result.InsertedID), nil
}
