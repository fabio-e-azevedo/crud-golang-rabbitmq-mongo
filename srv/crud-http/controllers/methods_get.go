package controllers

import (
	"crud-golang-rabbitmq-mongo/pkg/config"
	jph "crud-golang-rabbitmq-mongo/pkg/jsonplaceholder"
	"crud-golang-rabbitmq-mongo/pkg/mongodb"
	"crud-golang-rabbitmq-mongo/pkg/utils"

	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func GetAll(c *gin.Context) {
	resources := jph.NewResources()

	resourceType := strings.Split(c.Request.URL.Path, "/")[1]
	cfg := config.NewConfigMongo()
	m := mongodb.DbConnect{
		URI:        cfg.MongoURI,
		Database:   cfg.MongoDatabase,
		Collection: resourceType,
	}

	err := mongodb.FindAll(&resources, &m)
	utils.FailOnError(err, "Error finding all documents in mongo!")
	c.IndentedJSON(http.StatusOK, &resources)
}

func GetByID(c *gin.Context) {
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
	utils.FailOnError(err, "Error finding document in mongo!")
	c.IndentedJSON(http.StatusOK, &resource)
}
