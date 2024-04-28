package jsonplaceholder

import (
	"crud-golang-rabbitmq-mongo/internal"
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

func GetResources(data []byte) ([]IResource, error) {
	var result []IResource

	resultUnmarshal := []Resource{}
	err := newResources(&resultUnmarshal, data)
	if err != nil {
		return nil, err
	}

	for i := range resultUnmarshal {
		result = append(result, &resultUnmarshal[i])
	}
	return result, nil
}

func GetResource(data []byte) (IResource, error) {
	resultUnmarshal := Resource{}
	err := newResources(&resultUnmarshal, data)
	if err != nil {
		return nil, err
	}
	return &resultUnmarshal, nil
}

func newResources[T any](resource *T, data []byte) error {
	err := json.Unmarshal(data, &resource)
	internal.FailOnError(err, "Failed to Unmarshal newResources")
	return err
}
func NewResources() []Resource {
	return []Resource{}
}

func NewResource() IResource {
	return &Resource{}
}

func (p *Resource) Show() []byte {
	bodyResult, _ := json.Marshal(p)
	return bodyResult
}

func (p *ResourceMongo) Show() []byte {
	bodyResult, _ := json.Marshal(p)
	return bodyResult
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

	bodyResult, err := GetResources(body)
	if err != nil {
		fmt.Println(err)
	}

	return bodyResult, nil
}
