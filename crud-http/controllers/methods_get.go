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

// func GetAll(ctx *gin.Context) {
// 	resourceType := strings.Split(ctx.Request.URL.Path, "/")[3]

// 	cfg := config.NewConfigMongo()
// 	cfgMongo := mongodb.DbConnect{
// 		URI:        cfg.MongoURI,
// 		Database:   cfg.MongoDatabase,
// 		Collection: resourceType,
// 	}

// 	docBytes, docTotal, err := mongodb.FindAll(&cfgMongo)
// 	if err != nil {
// 		ctx.JSON(http.StatusNotFound, gin.H{
// 			"error": "not found documents in mongo",
// 		})
// 		return
// 	}

// 	result, err := jph.GetResources(resourceType, docTotal, docBytes)
// 	if err != nil {
// 		ctx.JSON(http.StatusForbidden, gin.H{
// 			"error": err,
// 		})
// 		return
// 	}

// 	ctx.JSON(http.StatusOK, &result)
// }

func getAll(resourceType string) ([]jph.IResource, error) {
	var resultBytes []byte
	var totalDocuments int
	var err error

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

// ListAlbums   godoc
// @Summary     List All Albums
// @Description get all albums
// @Tags        albums
// @Accept      json
// @Produce     json
// @Success     200  {array}   model.Album
// @Failure     404  {string}  "error"
// @Router      /albums [get]
func ListAlbums(ctx *gin.Context) {
	resourceType := strings.Split(ctx.Request.URL.Path, "/")[3]

	result, err := getAll(resourceType)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": err,
		})
		return
	}

	ctx.JSON(http.StatusOK, &result)
}

// ListComments godoc
// @Summary     List All Comments
// @Description get all comments
// @Tags        comments
// @Accept      json
// @Produce     json
// @Success     200  {array}   model.Comment
// @Failure     404  {string}  "error"
// @Router      /comments [get]
func ListComments(ctx *gin.Context) {
	resourceType := strings.Split(ctx.Request.URL.Path, "/")[3]

	result, err := getAll(resourceType)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": err,
		})
		return
	}

	ctx.JSON(http.StatusOK, &result)
}

// ListPhotos   godoc
// @Summary     List All Photos
// @Description get all photos
// @Tags        photos
// @Accept      json
// @Produce     json
// @Success     200  {array}   model.Photo
// @Failure     404  {string}  "error"
// @Router      /photos [get]
func ListPhotos(ctx *gin.Context) {
	resourceType := strings.Split(ctx.Request.URL.Path, "/")[3]

	result, err := getAll(resourceType)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": err,
		})
		return
	}

	ctx.JSON(http.StatusOK, &result)
}

// ListPosts    godoc
// @Summary     List All Posts
// @Description get all posts
// @Tags        posts
// @Accept      json
// @Produce     json
// @Success     200  {array}   model.Post
// @Failure     404  {string}  "error"
// @Router      /posts [get]
func ListPosts(ctx *gin.Context) {
	resourceType := strings.Split(ctx.Request.URL.Path, "/")[3]

	result, err := getAll(resourceType)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": err,
		})
		return
	}

	ctx.JSON(http.StatusOK, &result)
}

// ListTodos    godoc
// @Summary     List All Todos
// @Description get all todos
// @Tags        todos
// @Accept      json
// @Produce     json
// @Success     200  {array}   model.Todo
// @Failure     404  {string}  "error"
// @Router      /todos [get]
func ListTodos(ctx *gin.Context) {
	resourceType := strings.Split(ctx.Request.URL.Path, "/")[3]

	result, err := getAll(resourceType)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": err,
		})
		return
	}

	ctx.JSON(http.StatusOK, &result)
}

// ListUsers    godoc
// @Summary     List All Users
// @Description get all users
// @Tags        users
// @Accept      json
// @Produce     json
// @Success     200  {array}   model.User
// @Failure     404  {string}  "error"
// @Router      /users [get]
func ListUsers(ctx *gin.Context) {
	resourceType := strings.Split(ctx.Request.URL.Path, "/")[3]

	result, err := getAll(resourceType)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": err,
		})
		return
	}

	ctx.JSON(http.StatusOK, &result)
}

// func GetAllv1(configMongo *mongodb.DbConnect) gin.HandlerFunc {
// 	return func(ctx *gin.Context) {
// 		resources := jph.NewResources()

// 		resourceType := strings.Split(ctx.Request.URL.Path, "/")[3]

// 		cfgMongo := mongodb.DbConnect{
// 			Database:   configMongo.Database,
// 			Collection: resourceType,
// 		}

// 		err := mongodb.FindAll(&resources, &cfgMongo)
// 		if err != nil {
// 			ctx.JSON(http.StatusNotFound, gin.H{
// 				"error": "not found documents in mongo",
// 			})
// 			return
// 		}

// 		ctx.JSON(http.StatusOK, &resources)
// 	}
// }

func getByID(resourceType string, id int) (*jph.IResource, error) {
	cfg := config.NewConfigMongo()
	cfgMongo := mongodb.DbConnect{
		URI:        cfg.MongoURI,
		Database:   cfg.MongoDatabase,
		Collection: resourceType,
	}

	resultBytes, err := mongodb.FindOneById(id, &cfgMongo)
	if err != nil {
		return nil, fmt.Errorf("not found document in mongo")
	}

	resource, _ := jph.GetResource(resourceType, resultBytes)

	return &resource, nil
}

// ShowAlbum    godoc
// @Summary     Get Album By ID
// @Description get album by id
// @Tags        albums
// @Accept      json
// @Produce     json
// @Param       id   path      int  true  "Album ID" minimum(1)
// @Success     200  {object}  model.Album
// @Failure     404  {object}  httpError
// @Router      /albums/{id} [get]
func ShowAlbum(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		panic(err)
	}

	resourceType := strings.Split(ctx.Request.URL.Path, "/")[3]

	result, err := getByID(resourceType, id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, &httpError{Error: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, result)
}

// ShowComment  godoc
// @Summary     Get Comment By ID
// @Description get comment by id
// @Tags        comments
// @Accept      json
// @Produce     json
// @Param       id   path      int  true  "Comment ID" minimum(1)
// @Success     200  {object}  model.Comment
// @Failure     404  {object}  httpError
// @Router      /comments/{id} [get]
func ShowComment(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		panic(err)
	}

	resourceType := strings.Split(ctx.Request.URL.Path, "/")[3]

	result, err := getByID(resourceType, id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, &httpError{Error: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, result)
}

// ShowPhoto    godoc
// @Summary     Get Photo By ID
// @Description get photo by id
// @Tags        photos
// @Accept      json
// @Produce     json
// @Param       id   path      int  true  "Photo ID" minimum(1)
// @Success     200  {object}  model.Photo
// @Failure     404  {object}  httpError
// @Router      /photos/{id} [get]
func ShowPhoto(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		panic(err)
	}

	resourceType := strings.Split(ctx.Request.URL.Path, "/")[3]

	result, err := getByID(resourceType, id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, &httpError{Error: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, result)
}

// ShowPost     godoc
// @Summary     Get Post By ID
// @Description get post by id
// @Tags        posts
// @Accept      json
// @Produce     json
// @Param       id   path      int  true  "Post ID" minimum(1)
// @Success     200  {object}  model.Post
// @Failure     404  {object}  httpError
// @Router      /posts/{id} [get]
func ShowPost(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		panic(err)
	}

	resourceType := strings.Split(ctx.Request.URL.Path, "/")[3]

	result, err := getByID(resourceType, id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, &httpError{Error: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, result)
}

// ShowTodo     godoc
// @Summary     Get Todo By ID
// @Description get todo by id
// @Tags        todos
// @Accept      json
// @Produce     json
// @Param       id   path      int  true  "Todo ID" minimum(1)
// @Success     200  {object}  model.Todo
// @Failure     404  {object}  httpError
// @Router      /todos/{id} [get]
func ShowTodo(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		panic(err)
	}

	resourceType := strings.Split(ctx.Request.URL.Path, "/")[3]

	result, err := getByID(resourceType, id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, &httpError{Error: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, result)
}

// ShowUser     godoc
// @Summary     Get User By ID
// @Description get user by id
// @Tags        users
// @Accept      json
// @Produce     json
// @Param       id   path      int  true  "User ID" minimum(1)
// @Success     200  {object}  model.User
// @Failure     404  {object}  httpError
// @Router      /users/{id} [get]
func ShowUser(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		panic(err)
	}

	resourceType := strings.Split(ctx.Request.URL.Path, "/")[3]

	result, err := getByID(resourceType, id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, &httpError{Error: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, result)
}
