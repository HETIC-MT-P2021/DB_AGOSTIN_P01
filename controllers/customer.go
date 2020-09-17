package controllers

import (
	"github.com/gin-gonic/gin"
	"goevent/database"
	"goevent/database/Models"
	"log"
	"net/http"
)

func CustomerDetailsAction(c *gin.Context) {
	repository := Models.Repository{Conn: database.DbConn}
	customers, _ := repository.GetUser()

	log.Println(customers)

	c.JSON(http.StatusOK, customers)
}
