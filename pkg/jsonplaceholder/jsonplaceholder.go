package jsonplaceholder

import (
	"crud-golang-rabbitmq-mongo/pkg/model"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type IResource interface {
	Show() []byte
	Echo() string
	IsZeroId() bool
	SetId(n int)
}

func GetResources(resourceType string, length int, data []byte) ([]IResource, error) {
	var result []IResource
	var err error

	switch resourceType {
	case "albums":
		resources := make([]*model.Album, length)
		result, err = newResources(resources, data)
	case "comments":
		resources := make([]*model.Comment, length)
		result, err = newResources(resources, data)
	case "photos":
		resources := make([]*model.Photo, length)
		result, err = newResources(resources, data)
	case "posts":
		resources := make([]*model.Post, length)
		result, err = newResources(resources, data)
	case "todos":
		resources := make([]*model.Todo, length)
		result, err = newResources(resources, data)
	case "users":
		resources := make([]*model.User, length)
		result, err = newResources(resources, data)
	default:
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return result, err
}

func newResources[T IResource](resources []T, data []byte) ([]IResource, error) {
	err := json.Unmarshal(data, &resources)
	if err != nil {
		return nil, err
	}

	result := make([]IResource, len(resources))

	for i := range resources {
		result[i] = resources[i]
	}
	return result, nil
}

func GetResource(resourceType string, data []byte) (IResource, error) {
	resource, err := newResource(resourceType, data)
	if err != nil {
		return nil, err
	}
	return resource, nil
}

func newResource(resourceType string, data []byte) (IResource, error) {
	var resource IResource

	switch resourceType {
	case "albums":
		resource = &model.Album{}
	case "comments":
		resource = &model.Comment{}
	case "photos":
		resource = &model.Photo{}
	case "posts":
		resource = new(model.Post)
	case "todos":
		resource = new(model.Todo)
	case "users":
		resource = new(model.User)
	}

	err := json.Unmarshal(data, &resource)
	if err != nil {
		return nil, err
	}
	return resource, nil
}

func Get(url string, resource string, all bool) ([]IResource, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == 404 {
		return nil, fmt.Errorf("%s", string(body))
	}

	var bodyResult []IResource

	if all {
		bodyResult, err = GetResources(resource, 1, body)
		if err != nil {
			return nil, err
		}
	} else {
		result, err := GetResource(resource, body)
		if err != nil {
			return nil, err
		}

		bodyResult = append(bodyResult, result)
	}
	return bodyResult, nil
}
