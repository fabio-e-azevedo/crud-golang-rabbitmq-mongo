package controllers

import (
	"crud-golang-rabbitmq-mongo/pkg/config"
	"crud-golang-rabbitmq-mongo/pkg/jsonplaceholder"
	"crud-golang-rabbitmq-mongo/pkg/model"
	"crud-golang-rabbitmq-mongo/pkg/rabbitmq"

	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func postAll(resourceType string, content jsonplaceholder.IResource) {
	cfg := config.NewConfigRabbit()
	cfgRabbit := rabbitmq.RabbitMQ{
		URI:       cfg.RabbitURI,
		QueueName: resourceType,
	}
	cfgRabbit.Publisher(content.Show())
}

// AddAlbum     godoc
// @Summary     Post Create Album
// @Description post create album
// @Tags        albums
// @Accept      json
// @Produce     json
// @Param       request body model.Album true "album model for creation"
// @Success     204  {object}  model.Album
// @Failure     400  {object}  httpError
// @Router      /v1/albums [post]
func AddAlbum(ctx *gin.Context) {
	resourceType := strings.Split(ctx.Request.URL.Path, "/")[3]

	content := &model.Album{}
	if err := ctx.ShouldBindJSON(&content); err != nil {
		ctx.JSON(http.StatusBadRequest, &httpError{Error: err.Error()})
		return
	}

	postAll(resourceType, content)

	ctx.JSON(http.StatusOK, content)
}

// AddComment   godoc
// @Summary     Post Create Comment
// @Description post create comment
// @Tags        comments
// @Accept      json
// @Produce     json
// @Param       request body model.Comment true "comment model for creation"
// @Success     204  {object}  model.Comment
// @Failure     400  {object}  httpError
// @Router      /v1/comments [post]
func AddComment(ctx *gin.Context) {
	resourceType := strings.Split(ctx.Request.URL.Path, "/")[3]

	content := &model.Comment{}
	if err := ctx.ShouldBindJSON(&content); err != nil {
		ctx.JSON(http.StatusBadRequest, &httpError{Error: err.Error()})
		return
	}

	postAll(resourceType, content)

	ctx.JSON(http.StatusOK, content)
}

// AddPhoto     godoc
// @Summary     Post Create Photo
// @Description post create photo
// @Tags        photos
// @Accept      json
// @Produce     json
// @Param       request body model.Photo true "photo model for creation"
// @Success     204  {object}  model.Photo
// @Failure     400  {object}  httpError
// @Router      /v1/photos [post]
func AddPhoto(ctx *gin.Context) {
	resourceType := strings.Split(ctx.Request.URL.Path, "/")[3]

	content := &model.Photo{}
	if err := ctx.ShouldBindJSON(&content); err != nil {
		ctx.JSON(http.StatusBadRequest, &httpError{Error: err.Error()})
		return
	}

	postAll(resourceType, content)

	ctx.JSON(http.StatusOK, content)
}

// AddPost      godoc
// @Summary     Post Create Post
// @Description post create post
// @Tags        posts
// @Accept      json
// @Produce     json
// @Param       request body model.Post true "post model for creation"
// @Success     204  {object}  model.Post
// @Failure     400  {object}  httpError
// @Router      /v1/posts [post]
func AddPost(ctx *gin.Context) {
	resourceType := strings.Split(ctx.Request.URL.Path, "/")[3]

	content := &model.Post{}
	if err := ctx.ShouldBindJSON(&content); err != nil {
		ctx.JSON(http.StatusBadRequest, &httpError{Error: err.Error()})
		return
	}

	postAll(resourceType, content)

	ctx.JSON(http.StatusOK, content)
}

// AddTodo      godoc
// @Summary     Post Create Todo
// @Description post create todo
// @Tags        todos
// @Accept      json
// @Produce     json
// @Param       request body model.Todo true "todo model for creation"
// @Success     204  {object}  model.Todo
// @Failure     400  {object}  httpError
// @Router      /v1/todos [post]
func AddTodo(ctx *gin.Context) {
	resourceType := strings.Split(ctx.Request.URL.Path, "/")[3]

	content := &model.Todo{}
	if err := ctx.ShouldBindJSON(&content); err != nil {
		ctx.JSON(http.StatusBadRequest, &httpError{Error: err.Error()})
		return
	}

	postAll(resourceType, content)

	ctx.JSON(http.StatusOK, content)
}

// AddUser      godoc
// @Summary     Post Create User
// @Description post create user
// @Tags        users
// @Accept      json
// @Produce     json
// @Param       request body model.User true "user model for creation"
// @Success     204  {object}  model.User
// @Failure     400  {object}  httpError
// @Router      /v1/users [post]
func AddUser(ctx *gin.Context) {
	resourceType := strings.Split(ctx.Request.URL.Path, "/")[3]

	content := &model.User{}
	if err := ctx.ShouldBindJSON(&content); err != nil {
		ctx.JSON(http.StatusBadRequest, &httpError{Error: err.Error()})
		return
	}

	postAll(resourceType, content)

	ctx.JSON(http.StatusOK, content)
}
