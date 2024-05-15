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

func deleteByID(r string, id int) (string, error) {
	resourceType := r

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
// @Param       id   path      int  true  "Album ID" minimum(1)
// @Success     204  {object}  httpSuccess
// @Failure     404  {object}  httpError
// @Router      /albums/{id} [delete]
func DeleteAlbum(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		panic(err)
	}

	resourceType := strings.Split(c.Request.URL.Path, "/")[3]

	result, err := deleteByID(resourceType, id)
	if err != nil {
		c.JSON(http.StatusNotFound, &httpError{Error: err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, &httpSuccess{Result: result})
}

// DeleteComment godoc
// @Summary      Delete Comment By ID
// @Description  delete comment by id
// @Tags         comments
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Comment ID" minimum(1)
// @Success      204  {object}  httpSuccess
// @Failure      404  {object}  httpError
// @Router       /comments/{id} [delete]
func DeleteComment(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		panic(err)
	}

	resourceType := strings.Split(c.Request.URL.Path, "/")[3]

	result, err := deleteByID(resourceType, id)
	if err != nil {
		c.JSON(http.StatusNotFound, &httpError{Error: err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, &httpSuccess{Result: result})
}

// DeletePhoto  godoc
// @Summary     Delete Photo By ID
// @Description delete photo by id
// @Tags        photos
// @Accept      json
// @Produce     json
// @Param       id   path      int  true  "Photo ID" minimum(1)
// @Success     204  {object}  httpSuccess
// @Failure     404  {object}  httpError
// @Router      /photos/{id} [delete]
func DeletePhoto(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		panic(err)
	}

	resourceType := strings.Split(c.Request.URL.Path, "/")[3]

	result, err := deleteByID(resourceType, id)
	if err != nil {
		c.JSON(http.StatusNotFound, &httpError{Error: err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, &httpSuccess{Result: result})
}

// DeletePost   godoc
// @Summary     Delete Post By ID
// @Description delete post by id
// @Tags        posts
// @Accept      json
// @Produce     json
// @Param       id   path      int  true  "Post ID" minimum(1)
// @Success     204  {object}  httpSuccess
// @Failure     404  {object}  httpError
// @Router      /posts/{id} [delete]
func DeletePost(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		panic(err)
	}

	resourceType := strings.Split(c.Request.URL.Path, "/")[3]

	result, err := deleteByID(resourceType, id)
	if err != nil {
		c.JSON(http.StatusNotFound, &httpError{Error: err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, &httpSuccess{Result: result})
}

// DeleteTodo   godoc
// @Summary     Delete Todo By ID
// @Description delete todo by id
// @Tags        todos
// @Accept      json
// @Produce     json
// @Param       id   path      int  true  "Todo ID" minimum(1)
// @Success     204  {object}  httpSuccess
// @Failure     404  {object}  httpError
// @Router      /todos/{id} [delete]
func DeleteTodo(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		panic(err)
	}

	resourceType := strings.Split(c.Request.URL.Path, "/")[3]

	result, err := deleteByID(resourceType, id)
	if err != nil {
		c.JSON(http.StatusNotFound, &httpError{Error: err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, &httpSuccess{Result: result})
}

// DeleteUser   godoc
// @Summary     Delete User By ID
// @Description delete user by id
// @Tags        users
// @Accept      json
// @Produce     json
// @Param       id   path      int  true  "User ID" minimum(1)
// @Success     204  {object}  httpSuccess
// @Failure     404  {object}  httpError
// @Router      /users/{id} [delete]
func DeleteUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		panic(err)
	}

	resourceType := strings.Split(c.Request.URL.Path, "/")[3]

	result, err := deleteByID(resourceType, id)
	if err != nil {
		c.JSON(http.StatusNotFound, &httpError{Error: err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, &httpSuccess{Result: result})
}
