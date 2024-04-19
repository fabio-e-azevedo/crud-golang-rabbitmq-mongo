package main

import (
	"crud-golang-rabbitmq-mongo/internal"
	"crud-golang-rabbitmq-mongo/mongodb"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.GET("/users", getUsers)
	router.GET("/users/:id", getUserByID)
	//router.POST("/users", postUser)
	//router.PATCH("/users/:id", patchUserByID)
	//router.DELETE("/users/:id", deleteUserByID)
	router.Run("0.0.0.0:5000")
}

func getUsers(c *gin.Context) {
	result, err := mongodb.FindAll()
	internal.FailOnError(err, "Error finding all documents in mongo!")
	c.IndentedJSON(http.StatusOK, *result)
}

func getUserByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		panic(err)
	}

	result, err := mongodb.FindOne("id", id)
	internal.FailOnError(err, "Error finding document in mongo!")

	c.IndentedJSON(http.StatusOK, *result)
	//return
}
