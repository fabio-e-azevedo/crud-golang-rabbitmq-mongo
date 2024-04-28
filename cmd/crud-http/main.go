package main

import (
	"crud-golang-rabbitmq-mongo/internal"
	"crud-golang-rabbitmq-mongo/mongodb"
	"net/http"
	"strconv"
	"strings"

	"crud-golang-rabbitmq-mongo/config"
	jph "crud-golang-rabbitmq-mongo/jsonplaceholder"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.GET("/users", getAll)
	router.GET("/photos", getAll)
	router.GET("/posts", getAll)
	router.GET("/comments", getAll)
	router.GET("/albums", getAll)
	router.GET("/todos", getAll)

	router.GET("/albums/:id", getByID)
	router.GET("/todos/:id", getByID)
	router.GET("/users/:id", getByID)
	router.GET("/photos/:id", getByID)
	router.GET("/posts/:id", getByID)
	router.GET("/comments/:id", getByID)

	//router.POST("/users", postUser)
	//router.PATCH("/users/:id", patchUserByID)
	//router.DELETE("/users/:id", deleteUserByID)
	router.Run("0.0.0.0:5000")
}

func getAll(c *gin.Context) {
	resources := jph.NewResources()

	resourceType := strings.Split(c.Request.URL.Path, "/")[1]
	cfg := config.NewConfigMongo()
	m := mongodb.DbConnect{
		URI:        cfg.MongoURI,
		Database:   cfg.MongoDatabase,
		Collection: resourceType,
	}

	err := mongodb.FindAll(&resources, &m)
	internal.FailOnError(err, "Error finding all documents in mongo!")
	c.IndentedJSON(http.StatusOK, &resources)
}

func getByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		panic(err)
	}

	resourceType := strings.Split(c.Request.URL.Path, "/")[1]

	cfg := config.NewConfigMongo()
	m := mongodb.DbConnect{
		URI:        cfg.MongoURI,
		Database:   cfg.MongoDatabase,
		Collection: resourceType,
	}

	resource := jph.NewResource()

	err = mongodb.FindOne(&resource, "id", id, &m)
	internal.FailOnError(err, "Error finding document in mongo!")
	c.IndentedJSON(http.StatusOK, &resource)
}
