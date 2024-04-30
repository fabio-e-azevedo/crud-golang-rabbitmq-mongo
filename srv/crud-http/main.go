package main

import (
	"crud-golang-rabbitmq-mongo/srv/crud-http/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.GET("/users", controllers.GetAll)
	router.GET("/photos", controllers.GetAll)
	router.GET("/posts", controllers.GetAll)
	router.GET("/comments", controllers.GetAll)
	router.GET("/albums", controllers.GetAll)
	router.GET("/todos", controllers.GetAll)

	router.GET("/albums/:id", controllers.GetByID)
	router.GET("/todos/:id", controllers.GetByID)
	router.GET("/users/:id", controllers.GetByID)
	router.GET("/photos/:id", controllers.GetByID)
	router.GET("/posts/:id", controllers.GetByID)
	router.GET("/comments/:id", controllers.GetByID)

	//router.POST("/users", postUser)
	//router.PATCH("/users/:id", patchUserByID)
	//router.DELETE("/users/:id", deleteUserByID)
	router.Run("0.0.0.0:5000")
}
