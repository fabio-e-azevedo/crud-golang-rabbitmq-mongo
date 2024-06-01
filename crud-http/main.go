package main

import (
	"crud-golang-rabbitmq-mongo/crud-http/controllers"
	"log"
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
// @host                      cao-veio:80
// @BasePath                  /api
// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func main() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	v1 := router.Group("/api/v1")
	{
		albums := v1.Group("/albums")
		{
			albums.GET(":id", controllers.ShowAlbum)
			albums.GET("", controllers.ListAlbums)
			albums.POST("", controllers.AddAlbum)
			albums.DELETE(":id", controllers.DeleteAlbum)
			albums.PUT(":id", controllers.UpdateAlbum)
		}

		comments := v1.Group("/comments")
		{
			comments.GET(":id", controllers.ShowComment)
			comments.GET("", controllers.ListComments)
			comments.POST("", controllers.AddComment)
			comments.DELETE(":id", controllers.DeleteComment)
			comments.PUT(":id", controllers.UpdateComment)
		}

		photos := v1.Group("/photos")
		{
			photos.GET(":id", controllers.ShowPhoto)
			photos.GET("", controllers.ListPhotos)
			photos.POST("", controllers.AddPhoto)
			photos.DELETE(":id", controllers.DeletePhoto)
			photos.PUT(":id", controllers.UpdatePhoto)
		}

		posts := v1.Group("/posts")
		{
			posts.GET(":id", controllers.ShowPost)
			posts.GET("", controllers.ListPosts)
			posts.POST("", controllers.AddPost)
			posts.DELETE(":id", controllers.DeletePost)
			posts.PUT(":id", controllers.UpdatePost)
		}

		todos := v1.Group("/todos")
		{
			todos.GET(":id", controllers.ShowTodo)
			todos.GET("", controllers.ListTodos)
			todos.POST("", controllers.AddTodo)
			todos.DELETE(":id", controllers.DeleteTodo)
			todos.PUT(":id", controllers.UpdateTodo)
		}

		users := v1.Group("/users")
		{
			users.GET(":id", controllers.ShowUser)
			users.GET("", controllers.ListUsers)
			users.POST("", controllers.AddUser)
			users.DELETE(":id", controllers.DeleteUser)
			users.PUT(":id", controllers.UpdateUser)
		}
	}
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	router.GET("/swagger", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/swagger/index.html")
	})

	// v2 := router.Group("/api/v2")
	// v2.GET("/albums", controllers.GetAllv2)
	// v2.GET("/comments", controllers.GetAllv2)
	// v2.GET("/photos", controllers.GetAllv2)
	// v2.GET("/posts", controllers.GetAllv2)
	// v2.GET("/todos", controllers.GetAllv2)
	// v2.GET("/users", controllers.GetAllv2)

	//router.PATCH("/users/:id", patchUserByID)
	log.Fatal(router.Run("0.0.0.0:5000"))
}
