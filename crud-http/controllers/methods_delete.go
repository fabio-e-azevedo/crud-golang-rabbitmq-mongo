package controllers

import (
	"crud-golang-rabbitmq-mongo/pkg/config"
	"crud-golang-rabbitmq-mongo/pkg/mongodb"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type httpError struct {
	Error string `json:"error"`
}

type httpSuccess struct {
	Result string `json:"result"`
}

func deleteByID(resourceType string, id int) (string, error) {
	cfg := config.NewConfigMongo()
	cfgMongo := mongodb.DbConnect{
		URI:        cfg.MongoURI,
		Database:   cfg.MongoDatabase,
		Collection: resourceType,
	}

	err := mongodb.FindAndDelete(id, &cfgMongo)
	if err != nil {
		return "", fmt.Errorf(err.Error())
	}

	return fmt.Sprintf("successfully deleted document id %d\n", id), nil
}

// DeleteAlbum  godoc
// @Summary     Delete Album By ID
// @Description delete album by id
// @Tags        albums
// @Accept      json
// @Produce     json
// @Param       id   path      int  true  "Album ID"
// @Success     204  {object}  httpSuccess
// @Failure     404  {object}  httpError
// @Router      /v1/albums/{id} [delete]
func DeleteAlbum(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		panic(err)
	}

	resourceType := strings.Split(ctx.Request.URL.Path, "/")[3]

	result, err := deleteByID(resourceType, id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, &httpError{Error: err.Error()})
		return
	}

	ctx.JSON(http.StatusNoContent, &httpSuccess{Result: result})
}

// DeleteComment godoc
// @Summary      Delete Comment By ID
// @Description  delete comment by id
// @Tags         comments
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Comment ID"
// @Success      204  {object}  httpSuccess
// @Failure      404  {object}  httpError
// @Router       /v1/comments/{id} [delete]
func DeleteComment(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		panic(err)
	}

	resourceType := strings.Split(ctx.Request.URL.Path, "/")[3]

	result, err := deleteByID(resourceType, id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, &httpError{Error: err.Error()})
		return
	}

	ctx.JSON(http.StatusNoContent, &httpSuccess{Result: result})
}

// DeletePhoto  godoc
// @Summary     Delete Photo By ID
// @Description delete photo by id
// @Tags        photos
// @Accept      json
// @Produce     json
// @Param       id   path      int  true  "Photo ID"
// @Success     204  {object}  httpSuccess
// @Failure     404  {object}  httpError
// @Router      /v1/photos/{id} [delete]
func DeletePhoto(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		panic(err)
	}

	resourceType := strings.Split(ctx.Request.URL.Path, "/")[3]

	result, err := deleteByID(resourceType, id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, &httpError{Error: err.Error()})
		return
	}

	ctx.JSON(http.StatusNoContent, &httpSuccess{Result: result})
}

// DeletePost   godoc
// @Summary     Delete Post By ID
// @Description delete post by id
// @Tags        posts
// @Accept      json
// @Produce     json
// @Param       id   path      int  true  "Post ID"
// @Success     204  {object}  httpSuccess
// @Failure     404  {object}  httpError
// @Router      /v1/posts/{id} [delete]
func DeletePost(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		panic(err)
	}

	resourceType := strings.Split(ctx.Request.URL.Path, "/")[3]

	result, err := deleteByID(resourceType, id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, &httpError{Error: err.Error()})
		return
	}

	ctx.JSON(http.StatusNoContent, &httpSuccess{Result: result})
}

// DeleteTodo   godoc
// @Summary     Delete Todo By ID
// @Description delete todo by id
// @Tags        todos
// @Accept      json
// @Produce     json
// @Param       id   path      int  true  "Todo ID"
// @Success     204  {object}  httpSuccess
// @Failure     404  {object}  httpError
// @Router      /v1/todos/{id} [delete]
func DeleteTodo(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		panic(err)
	}

	resourceType := strings.Split(ctx.Request.URL.Path, "/")[3]

	result, err := deleteByID(resourceType, id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, &httpError{Error: err.Error()})
		return
	}

	ctx.JSON(http.StatusNoContent, &httpSuccess{Result: result})
}

// DeleteUser   godoc
// @Summary     Delete User By ID
// @Description delete user by id
// @Tags        users
// @Accept      json
// @Produce     json
// @Param       id   path      int  true  "User ID"
// @Success     204  {object}  httpSuccess
// @Failure     404  {object}  httpError
// @Router      /v1/users/{id} [delete]
func DeleteUser(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		panic(err)
	}

	resourceType := strings.Split(ctx.Request.URL.Path, "/")[3]

	result, err := deleteByID(resourceType, id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, &httpError{Error: err.Error()})
		return
	}

	ctx.JSON(http.StatusNoContent, &httpSuccess{Result: result})
}
