package controllers

import (
	"crud-golang-rabbitmq-mongo/pkg/config"
	jph "crud-golang-rabbitmq-mongo/pkg/jsonplaceholder"
	"crud-golang-rabbitmq-mongo/pkg/mongodb"
	"crud-golang-rabbitmq-mongo/pkg/rediscache"
	"log"

	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func GetAllv3(c *gin.Context) {
	resourceType := strings.Split(c.Request.URL.Path, "/")[3]

	cfg := config.NewConfigMongo()
	cfgMongo := mongodb.DbConnect{
		URI:        cfg.MongoURI,
		Database:   cfg.MongoDatabase,
		Collection: resourceType,
	}

	docBytes, docTotal, err := mongodb.FindAll(&cfgMongo)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "not found documents in mongo",
		})
		return
	}

	result, err := jph.GetResources(resourceType, docTotal, docBytes)
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusOK, &result)
}

func GetAll(c *gin.Context) {
	var resultBytes []byte
	var totalDocuments int
	var err error

	resourceType := strings.Split(c.Request.URL.Path, "/")[3]

	resultRedis, err := rediscache.Get(resourceType)
	if err != nil {
		log.Printf("REDIS: %v", err)
		cfg := config.NewConfigMongo()
		cfgMongo := mongodb.DbConnect{
			URI:        cfg.MongoURI,
			Database:   cfg.MongoDatabase,
			Collection: resourceType,
		}

		resultBytes, totalDocuments, err = mongodb.FindAll(&cfgMongo)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "not found documents in mongo",
			})
			return
		}

		rediscache.Set(resourceType, resultBytes)
	} else {
		// totalDocuments = 5000
		resultBytes = []byte(resultRedis)
	}

	result, err := jph.GetResources(resourceType, totalDocuments, resultBytes)
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{
			"error": err,
		})
		return
	}
	c.JSON(http.StatusOK, &result)
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
		c.JSON(http.StatusNotFound, gin.H{
			"error": "not found document in mongo",
		})
		return
	}

	c.JSON(http.StatusOK, &resource)
}
