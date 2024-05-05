package controllers

import (
	"crud-golang-rabbitmq-mongo/pkg/config"
	jph "crud-golang-rabbitmq-mongo/pkg/jsonplaceholder"
	"crud-golang-rabbitmq-mongo/pkg/mongodb"

	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func GetAll(c *gin.Context) {
	resources := jph.NewResources()

	resourceType := strings.Split(c.Request.URL.Path, "/")[3]

	cfg := config.NewConfigMongo()
	cfgMongo := mongodb.DbConnect{
		URI:        cfg.MongoURI,
		Database:   cfg.MongoDatabase,
		Collection: resourceType,
	}

	err := mongodb.FindAll(&resources, &cfgMongo)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "not found documents in mongo",
		})
		return
	}

	c.JSON(http.StatusOK, &resources)
}

// func GetAllv1(configMongo *mongodb.DbConnect) gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		resources := jph.NewResources()

// 		resourceType := strings.Split(c.Request.URL.Path, "/")[3]

// 		cfgMongo := mongodb.DbConnect{
// 			Database:   configMongo.Database,
// 			Collection: resourceType,
// 		}

// 		err := mongodb.FindAll(&resources, &cfgMongo)
// 		if err != nil {
// 			c.JSON(http.StatusNotFound, gin.H{
// 				"error": "not found documents in mongo",
// 			})
// 			return
// 		}

// 		c.JSON(http.StatusOK, &resources)
// 	}
// }

func GetByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		panic(err)
	}

	resourceType := strings.Split(c.Request.URL.Path, "/")[3]

	cfg := config.NewConfigMongo()
	cfgMongo := mongodb.DbConnect{
		URI:        cfg.MongoURI,
		Database:   cfg.MongoDatabase,
		Collection: resourceType,
	}

	resource := jph.NewResource()

	err = mongodb.FindOne(&resource, "id", id, &cfgMongo)
	if err != nil {
		//c.AbortWithStatus(http.StatusNotFound)
		c.JSON(http.StatusNotFound, gin.H{
			"error": "not found document in mongo",
		})
		return
	}

	c.JSON(http.StatusOK, &resource)
}
