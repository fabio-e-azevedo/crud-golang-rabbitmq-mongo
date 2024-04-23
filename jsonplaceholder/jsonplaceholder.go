package jsonplaceholder

import (
	"crud-golang-rabbitmq-mongo/internal"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type ResourceGeneric interface {
	User | Photo | Posts
}

type GenericResource[T any] interface {
	New(data []byte) T
}

type Resource[T any] struct {
	ResourceType string `json:"resourcetype"`
	Data         []T    `json:"data"`
}

type Posts struct {
	UserId int32  `json:"userId" bson:"userId"`
	Id     int32  `json:"id" bson:"id"`
	Title  string `json:"title" bson:"title"`
	Body   string `json:"body" bson:"body"`
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

func (u *User) New(data []byte) User {
	err := json.Unmarshal(data, u)
	internal.FailOnError(err, "Failed to Unmarshal User")
	return *u
}

func (p *Photo) New(data []byte) Photo {
	err := json.Unmarshal(data, p)
	internal.FailOnError(err, "Failed to Unmarshal Photo")
	return *p
}

func (p *Posts) New(data []byte) Posts {
	err := json.Unmarshal(data, p)
	internal.FailOnError(err, "Failed to Unmarshal Posts")
	return *p
}

func (it *Resource[T]) New(body []byte) Resource[T] {
	err := json.Unmarshal(body, it)
	internal.FailOnError(err, "Failed to Unmarshal Resource[T]")
	return *it
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
	//return body, nil
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
	case "posts":
		bodyResult, err = getJson[Posts](resource, body)
		if err != nil {
			return nil, err
		}
	}

	return bodyResult, nil
}
