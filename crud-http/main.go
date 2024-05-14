package main

import (
	"crud-golang-rabbitmq-mongo/crud-http/controllers"
	"net/http"

	_ "crud-golang-rabbitmq-mongo/pkg/docs"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title                     CRUD API REST
// @version                   1.0.0
// @termsOfService            http://swagger.io/terms/
// @contact.name              Fabio Azevedo
// @description               API REST created to apply knowledge in learning the Go language
// @license.name              MIT
// @license.url               https://www.mit.edu/~amini/LICENSE.md
// @host                      cao-veio:5000
// @BasePath                  /api/v1
// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func main() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	v1 := router.Group("/api/v1")
	v1.GET("/albums", controllers.GetAlbums)
	v1.GET("/comments", controllers.GetComments)
	v1.GET("/photos", controllers.GetPhotos)
	v1.GET("/posts", controllers.GetPosts)
	v1.GET("/todos", controllers.GetTodos)
	v1.GET("/users", controllers.GetUsers)

	v1.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	v1.GET("/swagger", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/api/v1/swagger/index.html")
	})

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

	// v2 := router.Group("/api/v2")
	// v2.GET("/albums", controllers.GetAllv2)
	// v2.GET("/comments", controllers.GetAllv2)
	// v2.GET("/photos", controllers.GetAllv2)
	// v2.GET("/posts", controllers.GetAllv2)
	// v2.GET("/todos", controllers.GetAllv2)
	// v2.GET("/users", controllers.GetAllv2)

	//router.PATCH("/users/:id", patchUserByID)
	//router.DELETE("/users/:id", deleteUserByID)
	router.Run("0.0.0.0:5000")
}
