package main

import (
	"crud-golang-rabbitmq-mongo/pkg/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	v1 := router.Group("/api/v1")
	v1.GET("/albums", controllers.GetAll)
	v1.GET("/comments", controllers.GetAll)
	v1.GET("/photos", controllers.GetAll)
	v1.GET("/posts", controllers.GetAll)
	v1.GET("/todos", controllers.GetAll)
	v1.GET("/users", controllers.GetAll)

	v1.GET("/albums/:id", controllers.GetByID)
	v1.GET("/comments/:id", controllers.GetByID)
	v1.GET("/photos/:id", controllers.GetByID)
	v1.GET("/posts/:id", controllers.GetByID)
	v1.GET("/todos/:id", controllers.GetByID)
	v1.GET("/users/:id", controllers.GetByID)

	v1.POST("/albums", controllers.PostAll)
	v1.POST("/comments", controllers.PostAll)
	v1.POST("/photos", controllers.PostAll)
	v1.POST("/posts", controllers.PostAll)
	v1.POST("/todos", controllers.PostAll)
	v1.POST("/users", controllers.PostAll)

	v1.DELETE("/albums/:id", controllers.DeleteByID)
	v1.DELETE("/comments/:id", controllers.DeleteByID)
	v1.DELETE("/photos/:id", controllers.DeleteByID)
	v1.DELETE("/posts/:id", controllers.DeleteByID)
	v1.DELETE("/todos/:id", controllers.DeleteByID)
	v1.DELETE("/users/:id", controllers.DeleteByID)

	//router.PATCH("/users/:id", patchUserByID)
	//router.DELETE("/users/:id", deleteUserByID)
	router.Run("0.0.0.0:5000")
}
