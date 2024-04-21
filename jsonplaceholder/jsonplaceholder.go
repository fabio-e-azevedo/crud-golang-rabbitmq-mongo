package jsonplaceholder

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type ResourceGeneric interface {
	User | Photo
}

type Resource[T any] struct {
	ResourceType string `json:"resourcetype"`
	Data         []T    `json:"data"`
}

type Photo struct {
	AlbumId      int32  `json:"albumId" bson:"albumId"`
	Id           int32  `json:"id" bson:"id"`
	Title        string `json:"title" bson:"title"`
	Url          string `json:"url" bson:"url"`
	ThumbnailUrl string `json:"thumbnailUrl" bson:"thumbnailUrl"`
}

type User struct {
	Id       int32  `json:"id" bson:"id"`
	Name     string `json:"name" bson:"name"`
	Username string `json:"username" bson:"username"`
	Email    string `json:"email" bson:"email"`
	Address  struct {
		Street  string `json:"street" bson:"street"`
		Suite   string `json:"suite" bson:"suite"`
		City    string `json:"city" bson:"city"`
		ZipCode string `json:"zipcode" bson:"zipcode"`
		Geo     struct {
			Lat string `json:"lat" bson:"lat"`
			Lng string `json:"lng" bson:"lng"`
		}
	}
	Phone   string `json:"phone" bson:"phone"`
	Website string `json:"website" bson:"website"`
	Company struct {
		Name        string `json:"name" bson:"name"`
		CatchPhrase string `json:"catchPhrase" bson:"catchPhrase"`
		Bs          string `json:"bs" bson:"bs"`
	}
}

func getJson[T ResourceGeneric](resource string, body []byte) ([]byte, error) {
	var jsonTostruct []T
	var resources Resource[T]

	resources.ResourceType = resource

	err := json.Unmarshal(body, &jsonTostruct)
	if err != nil {
		return nil, err
	}

	resources.Data = jsonTostruct
	bodyResult, _ := json.Marshal(resources)

	return bodyResult, nil
}

func Get(resource string) ([]byte, error) {
	resp, err := http.Get(fmt.Sprintf("https://jsonplaceholder.typicode.com/%s", resource))
	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var bodyResult []byte
	switch resource {
	case "users":
		bodyResult, err = getJson[User](resource, body)
		if err != nil {
			return nil, err
		}
	case "photos":
		bodyResult, err = getJson[Photo](resource, body)
		if err != nil {
			return nil, err
		}
	}

	return bodyResult, nil
}
