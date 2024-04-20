package jsonplaceholder

import (
	"fmt"
	"io"
	"net/http"
)

type Resource struct {
	ResourceType string       `json:"resourcetype"`
	Data         UserFromJson `json:"data"`
}

type UserFromJson struct {
	Id       int32  `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Address  struct {
		Street  string `json:"street"`
		Suite   string `json:"suite"`
		City    string `json:"city"`
		ZipCode string `json:"zipcode"`
		Geo     struct {
			Lat string `json:"lat"`
			Lng string `json:"lng"`
		}
	}
	Phone   string `json:"phone"`
	Website string `json:"website"`
	Company struct {
		Name        string `json:"name"`
		CatchPhrase string `json:"catchPhrase"`
		Bs          string `json:"bs"`
	}
}

type UserFromBson struct {
	Id       int32  `bson:"id"`
	Name     string `bson:"name"`
	Username string `bson:"username"`
	Email    string `bson:"email"`
	Address  struct {
		Street  string `bson:"street"`
		Suite   string `bson:"suite"`
		City    string `bson:"city"`
		ZipCode string `bson:"zipcode"`
		Geo     struct {
			Lat string `bson:"lat"`
			Lng string `bson:"lng"`
		}
	}
	Phone   string `bson:"phone"`
	Website string `bson:"website"`
	Company struct {
		Name        string `bson:"name"`
		CatchPhrase string `bson:"catchPhrase"`
		Bs          string `bson:"bs"`
	}
}

func Get(resource string) (*[]byte, error) {
	resp, err := http.Get(fmt.Sprintf("https://jsonplaceholder.typicode.com/%s", resource))
	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return &body, nil

	// var users []UserFromJson

	// err = json.Unmarshal(body, &users)
	// if err != nil {
	// 	return nil, err
	// }

	// result := User{
	// 	Resource: "users",
	// 	Data:     users,
	// }

	// return &result, nil
}
