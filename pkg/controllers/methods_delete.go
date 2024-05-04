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

func DeleteByID(c *gin.Context) {
	idNumber, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		panic(err)
	}

	resourceType := strings.Split(c.Request.URL.Path, "/")[3]

	cfg := config.NewConfigMongo()
	m := mongodb.DbConnect{
		URI:        cfg.MongoURI,
		Database:   cfg.MongoDatabase,
		Collection: resourceType,
	}

	err = mongodb.FindAndDelete(idNumber, &m)
	if err != nil {
		c.JSON(http.StatusNotFound, fmt.Sprint(err))
		return
	}

	c.JSON(http.StatusNoContent, fmt.Sprintf("successfully deleted document id %d\n", idNumber))
}
