package cmd

import (
	"errors"
	"strings"
)

type validResourceValue string

var (
	validResourceOptions = []string{"albums", "comments", "photos", "posts", "todos", "users"}
)

func (v *validResourceValue) Set(value string) error {
	lowerValue := strings.ToLower(value)
	for _, option := range validResourceOptions {
		if lowerValue == option {
			*v = validResourceValue(lowerValue)
			return nil
		}
	}
	return errors.New("\nvalid options: albums, comments, photos, posts, todos, users")
}

func (v *validResourceValue) String() string {
	return string(*v)
}

func (v *validResourceValue) Type() string {
	return "name"
}
