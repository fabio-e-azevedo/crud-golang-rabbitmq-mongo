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
}

func GetResources(resourceType string, data []byte) ([]IResource, error) {
	var result []IResource
	var err error

	switch resourceType {
	case "albums":
		result, err = newAlbums(data)
	case "comments":
		result, err = newComments(data)
	case "photos":
		result, err = newPhotos(data)
	case "posts":
		result, err = newPosts(data)
	case "todos":
		result, err = newTodos(data)
	case "users":
		result, err = newUsers(data)
	}

	if err != nil {
		return nil, err
	}
	return result, err
}

func GetResource(resourceType string, data []byte) (IResource, error) {
	resource, err := newResource(resourceType, data)
	if err != nil {
		return nil, err
	}
	return resource, nil
}

// func newResources[T any](resource *T, data []byte) error {
// 	err := json.Unmarshal(data, &resource)
// 	internal.FailOnError(err, "Failed to Unmarshal newResources")
// 	return err
// }

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

func newAlbums(data []byte) ([]IResource, error) {
	var resources []Album
	err := json.Unmarshal(data, &resources)
	if err != nil {
		return nil, err
	}
	result := []IResource{}
	for i := range resources {
		result = append(result, &resources[i])
	}
	return result, nil
}

func newComments(data []byte) ([]IResource, error) {
	var resources []Comment
	err := json.Unmarshal(data, &resources)
	if err != nil {
		return nil, err
	}
	result := []IResource{}
	for i := range resources {
		result = append(result, &resources[i])
	}
	return result, nil
}

func newPhotos(data []byte) ([]IResource, error) {
	var resources []Photo
	err := json.Unmarshal(data, &resources)
	if err != nil {
		return nil, err
	}
	result := []IResource{}
	for i := range resources {
		result = append(result, &resources[i])
	}
	return result, nil
}

func newPosts(data []byte) ([]IResource, error) {
	var resources []Post
	err := json.Unmarshal(data, &resources)
	if err != nil {
		return nil, err
	}
	result := []IResource{}
	for i := range resources {
		result = append(result, &resources[i])
	}
	return result, nil
}

func newTodos(data []byte) ([]IResource, error) {
	var resources []Todo
	err := json.Unmarshal(data, &resources)
	if err != nil {
		return nil, err
	}
	result := []IResource{}
	for i := range resources {
		result = append(result, &resources[i])
	}
	return result, nil
}

func newUsers(data []byte) ([]IResource, error) {
	var resources []User
	err := json.Unmarshal(data, &resources)
	if err != nil {
		return nil, err
	}
	result := []IResource{}
	for i := range resources {
		result = append(result, &resources[i])
	}
	return result, nil
}

func NewResources() []Resource {
	return []Resource{}
}

func NewResource() IResource {
	return &Resource{}
}

func Get(resource string) ([]IResource, error) {
	resp, err := http.Get(fmt.Sprintf("https://jsonplaceholder.typicode.com/%s", resource))
	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	bodyResult, err := GetResources(resource, body)
	if err != nil {
		fmt.Println(err)
	}

	return bodyResult, nil
}
