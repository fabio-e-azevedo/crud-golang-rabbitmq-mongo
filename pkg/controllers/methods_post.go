package controllers

import (
	"crud-golang-rabbitmq-mongo/pkg/config"
	jph "crud-golang-rabbitmq-mongo/pkg/jsonplaceholder"
	"crud-golang-rabbitmq-mongo/pkg/rabbitmq"

	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func PostAll(ctx *gin.Context) {
	var content jph.IResource

	resourceType := strings.Split(ctx.Request.URL.Path, "/")[3]

	switch resourceType {
	case "albums":
		content = &jph.Album{}
	case "comments":
		content = &jph.Comment{}
	case "photos":
		content = &jph.Photo{}
	case "posts":
		content = &jph.Post{}
	case "todos":
		content = &jph.Todo{}
	case "users":
		content = &jph.User{}
	}

	if err := ctx.ShouldBindJSON(&content); err != nil {
		ctx.Error(err)
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	sendMessageRabbitMq(resourceType, content.Show())
	ctx.JSON(http.StatusOK, content)
}

func sendMessageRabbitMq(resourceType string, msg []byte) {
	cfg := config.NewConfigRabbit()
	rabbit := rabbitmq.RabbitMQ{
		URI:       cfg.RabbitURI,
		QueueName: resourceType,
	}

	rabbit.Publisher(msg)

}
