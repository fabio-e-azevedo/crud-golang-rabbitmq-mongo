package controllers

import (
	"crud-golang-rabbitmq-mongo/pkg/config"
	jph "crud-golang-rabbitmq-mongo/pkg/jsonplaceholder"
	"crud-golang-rabbitmq-mongo/pkg/mongodb"
	"crud-golang-rabbitmq-mongo/pkg/rediscache"
	"fmt"
	"log"

	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

// func GetAllv2(c *gin.Context) {
// 	resourceType := strings.Split(c.Request.URL.Path, "/")[3]

// 	cfg := config.NewConfigMongo()
// 	cfgMongo := mongodb.DbConnect{
// 		URI:        cfg.MongoURI,
// 		Database:   cfg.MongoDatabase,
// 		Collection: resourceType,
// 	}

// 	docBytes, docTotal, err := mongodb.FindAll(&cfgMongo)
// 	if err != nil {
// 		c.JSON(http.StatusNotFound, gin.H{
// 			"error": "not found documents in mongo",
// 		})
// 		return
// 	}

// 	result, err := jph.GetResources(resourceType, docTotal, docBytes)
// 	if err != nil {
// 		c.JSON(http.StatusForbidden, gin.H{
// 			"error": err,
// 		})
// 		return
// 	}

// 	c.JSON(http.StatusOK, &result)
// }

func getAllV2(r string) ([]jph.IResource, error) {
	var resultBytes []byte
	var totalDocuments int
	var err error

	resourceType := r

	resultRedis, err := rediscache.Get(resourceType)
	if err != nil {
		log.SetPrefix("[RDS] ")
		log.Printf("| %v", err)
		log.SetPrefix("")

		cfg := config.NewConfigMongo()
		cfgMongo := mongodb.DbConnect{
			URI:        cfg.MongoURI,
			Database:   cfg.MongoDatabase,
			Collection: resourceType,
		}

		resultBytes, totalDocuments, err = mongodb.FindAll(&cfgMongo)
		if err != nil {
			return nil, fmt.Errorf("not found documents in mongo") //http.StatusNotFound untyped int = 404
		}

		rediscache.Set(resourceType, resultBytes)
	} else {
		resultBytes = []byte(resultRedis)
	}

	result, err := jph.GetResources(resourceType, totalDocuments, resultBytes)
	if err != nil {
		return nil, fmt.Errorf(err.Error()) //http.StatusForbidden untyped int = 403
	}

	return result, nil
}

// ListAlbums godoc
// @Summary      List albums
// @Description  get all albums
// @Tags         albums
// @Accept       json
// @Produce      json
// @Success      200  {array}   jph.Album
// @Failure      404  {string}  "error"
// @Router       /albums [get]
func GetAlbums(c *gin.Context) {
	resourceType := strings.Split(c.Request.URL.Path, "/")[3]

	result, err := getAllV2(resourceType)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusOK, &result)
}

// ListComments godoc
// @Summary      List comments
// @Description  get all comments
// @Tags         comments
// @Accept       json
// @Produce      json
// @Success      200  {array}   jph.Comment
// @Failure      404  {string}  "error"
// @Router       /comments [get]
func GetComments(c *gin.Context) {
	resourceType := strings.Split(c.Request.URL.Path, "/")[3]

	result, err := getAllV2(resourceType)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusOK, &result)
}

// ListPhotos godoc
// @Summary      List photos
// @Description  get all photos
// @Tags         photos
// @Accept       json
// @Produce      json
// @Success      200  {array}   jph.Photo
// @Failure      404  {string}  "error"
// @Router       /photos [get]
func GetPhotos(c *gin.Context) {
	resourceType := strings.Split(c.Request.URL.Path, "/")[3]

	result, err := getAllV2(resourceType)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusOK, &result)
}

// ListPosts godoc
// @Summary      List posts
// @Description  get all posts
// @Tags         posts
// @Accept       json
// @Produce      json
// @Success      200  {array}   jph.Post
// @Failure      404  {string}  "error"
// @Router       /posts [get]
func GetPosts(c *gin.Context) {
	resourceType := strings.Split(c.Request.URL.Path, "/")[3]

	result, err := getAllV2(resourceType)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusOK, &result)
}

// ListTodos godoc
// @Summary      List todos
// @Description  get all todos
// @Tags         todos
// @Accept       json
// @Produce      json
// @Success      200  {array}   jph.Todo
// @Failure      404  {string}  "error"
// @Router       /todos [get]
func GetTodos(c *gin.Context) {
	resourceType := strings.Split(c.Request.URL.Path, "/")[3]

	result, err := getAllV2(resourceType)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusOK, &result)
}

// ListUsers godoc
// @Summary      List users
// @Description  get all users
// @Tags         users
// @Accept       json
// @Produce      json
// @Success      200  {array}   jph.User
// @Failure      404  {string}  "error"
// @Router       /users [get]
func GetUsers(c *gin.Context) {
	resourceType := strings.Split(c.Request.URL.Path, "/")[3]

	result, err := getAllV2(resourceType)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
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

	resultBytes, err := mongodb.FindOneById(id, &cfgMongo)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "not found document in mongo",
		})
		return
	}

	resource, _ := jph.GetResource(resourceType, resultBytes)

	c.JSON(http.StatusOK, &resource)
}
