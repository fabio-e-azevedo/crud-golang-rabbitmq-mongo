package controllers

import (
	"crud-golang-rabbitmq-mongo/pkg/jsonplaceholder"
	"crud-golang-rabbitmq-mongo/pkg/model"
	"log"
	"strconv"

	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func updateAll(ctx *gin.Context, resource *jsonplaceholder.IResource) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		panic(err)
	}
	//resourceType := strings.Split(ctx.Request.URL.Path, "/")[3]

	log.Printf("ID: %d\n", id)
	log.Println(*resource)

	err = ctx.BindJSON(resource)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, &httpError{Error: err.Error()})
		return
	}

	validate := validator.New()
	err = validate.Struct(*resource)
	if err != nil {
		validationErrors := err.(validator.ValidationErrors)
		ctx.JSON(http.StatusBadRequest, &httpError{Error: validationErrors.Error()})
		return
	}

	(*resource).SetId(id)
	log.Println(*resource)

	ctx.JSON(http.StatusOK, resource)
}

// UpdateAlbum  godoc
// @Summary     Update Album
// @Description update album
// @Tags        albums
// @Accept      json
// @Produce     json
// @Param       id   path      int  true  "Album ID"
// @Param       request body model.Album true "album model for update"
// @Success     204  {object}  model.Album
// @Failure     400  {object}  httpError
// @Router      /v1/albums [put]
func UpdateAlbum(ctx *gin.Context) {
	var resource jsonplaceholder.IResource = &model.Album{}
	updateAll(ctx, &resource)
}

// UpdateComment godoc
// @Summary      Update Comment
// @Description  update comment
// @Tags         comments
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Comment ID"
// @Param        request body model.Comment true "comment model for update"
// @Success      204  {object}  model.Comment
// @Failure      400  {object}  httpError
// @Router       /v1/comments [put]
func UpdateComment(ctx *gin.Context) {
	var resource jsonplaceholder.IResource = &model.Comment{}
	updateAll(ctx, &resource)
}

// UpdatePhoto  godoc
// @Summary     Update Photo
// @Description update photo
// @Tags        photos
// @Accept      json
// @Produce     json
// @Param       id   path      int  true  "Photo ID"
// @Param       request body model.Photo true "photo model for update"
// @Success     204  {object}  model.Photo
// @Failure     400  {object}  httpError
// @Router      /v1/photos [put]
func UpdatePhoto(ctx *gin.Context) {
	var resource jsonplaceholder.IResource = &model.Photo{}
	updateAll(ctx, &resource)
}

// UpdatePost   godoc
// @Summary     Update Post
// @Description update post
// @Tags        posts
// @Accept      json
// @Produce     json
// @Param       id   path      int  true  "Post ID"
// @Param       request body model.Post true "post model for update"
// @Success     204  {object}  model.Post
// @Failure     400  {object}  httpError
// @Router      /v1/posts [put]
func UpdatePost(ctx *gin.Context) {
	var resource jsonplaceholder.IResource = &model.Post{}
	updateAll(ctx, &resource)
}

// UpdateTodo   godoc
// @Summary     Update Todo
// @Description update todo
// @Tags        todos
// @Accept      json
// @Produce     json
// @Param       id   path      int  true  "Todo ID"
// @Param       request body model.Todo true "todo model for update"
// @Success     204  {object}  model.Todo
// @Failure     400  {object}  httpError
// @Router      /v1/todos [put]
func UpdateTodo(ctx *gin.Context) {
	var resource jsonplaceholder.IResource = &model.Todo{}
	updateAll(ctx, &resource)
}

// UpdateUser   godoc
// @Summary     Update User
// @Description update user
// @Tags        users
// @Accept      json
// @Produce     json
// @Param       id   path      int  true  "User ID"
// @Param       request body model.User true "user model for update"
// @Success     204  {object}  model.User
// @Failure     400  {object}  httpError
// @Router      /v1/users [put]
func UpdateUser(ctx *gin.Context) {
	var resource jsonplaceholder.IResource = &model.User{}
	updateAll(ctx, &resource)
}
