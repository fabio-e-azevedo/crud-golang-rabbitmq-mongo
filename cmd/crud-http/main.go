package main

import (
	"crud-golang-rabbitmq-mongo/internal"
	"crud-golang-rabbitmq-mongo/mongodb"
	"net/http"
	"strconv"
	"strings"

	jph "crud-golang-rabbitmq-mongo/jsonplaceholder"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.GET("/users", getAll)
	router.GET("/photos", getAll)
	router.GET("/posts", getAll)
	router.GET("/users/:id", getByID)
	router.GET("/photos/:id", getByID)
	router.GET("/posts/:id", getByID)
	//router.POST("/users", postUser)
	//router.PATCH("/users/:id", patchUserByID)
	//router.DELETE("/users/:id", deleteUserByID)
	router.Run("0.0.0.0:5000")
}

func getAll(c *gin.Context) {
	switch c.Request.URL.Path {
	case "/users":
		genericGetAll[jph.User]("users", c)
	case "/photos":
		genericGetAll[jph.Photo]("photos", c)
	case "/posts":
		genericGetAll[jph.Posts]("posts", c)
	}
}

func genericGetAll[T jph.ResourceGeneric](resourceEndpoint string, c *gin.Context) {
	resource := []T{}
	result, err := mongodb.FindAll(&resource, resourceEndpoint)
	internal.FailOnError(err, "Error finding all documents in mongo!")
	c.IndentedJSON(http.StatusOK, *result)
}

func getByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		panic(err)
	}

	switch strings.Split(c.Request.URL.Path, "/")[1] {
	case "users":
		genericGetByID[jph.User]("users", id, c)
	case "photos":
		genericGetByID[jph.Photo]("photos", id, c)
	case "posts":
		genericGetByID[jph.Posts]("posts", id, c)
	}
}

func genericGetByID[T jph.ResourceGeneric](resourceEndpoint string, id int, c *gin.Context) {
	var resource T
	result, err := mongodb.FindOne(&resource, "id", id, resourceEndpoint)
	internal.FailOnError(err, "Error finding document in mongo!")
	c.IndentedJSON(http.StatusOK, *result)
}
