package jsonplaceholder

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type ValidResource interface {
	User | Photo | Post | Comment | Album | Todo
}

type IResource interface {
	Show() []byte
	Echo() string
}

func GetResources(resourceType string, data []byte) ([]IResource, error) {
	var result []IResource
	var err error

	switch resourceType {
	case "albums":
		var resources []*Album
		result, err = newResources(resources, data)
	case "comments":
		var resources []*Comment
		result, err = newResources(resources, data)
	case "photos":
		var resources []*Photo
		result, err = newResources(resources, data)
	case "posts":
		var resources []*Post
		result, err = newResources(resources, data)
	case "todos":
		var resources []*Todo
		result, err = newResources(resources, data)
	case "users":
		var resources []*User
		result, err = newResources(resources, data)
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
	result := []IResource{}
	for i := range resources {
		result = append(result, resources[i])
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
		resource = &Album{}
	case "comments":
		resource = &Comment{}
	case "photos":
		resource = &Photo{}
	case "posts":
		resource = &Post{}
	case "todos":
		resource = &Todo{}
	case "users":
		resource = &User{}
	}
	err := json.Unmarshal(data, &resource)
	if err != nil {
		return nil, err
	}
	return resource, nil
}

func NewResources() []Resource {
	return []Resource{}
}

func NewResource() IResource {
	return &Resource{}
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

	var bodyResult []IResource

	if all {
		bodyResult, err = GetResources(resource, body)
		if err != nil {
			fmt.Println(err)
		}
	} else {
		result, err := GetResource(resource, body)
		if err != nil {
			fmt.Println(err)
		}

		bodyResult = append(bodyResult, result)
	}
	return bodyResult, nil
}
