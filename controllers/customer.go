package controllers

import (
	"github.com/gin-gonic/gin"
	"goevent/database"
	"goevent/database/Models"
	"net/http"
)

func AllCustomerAction(c *gin.Context) {
	repository := Models.Repository{Conn: database.DbConn}
	customers, _ := repository.GetAllCustomers()
	c.JSON(http.StatusOK, customers)
}

func OrderByCustomer(c *gin.Context) {
	repository := Models.Repository{Conn: database.DbConn}
	customers, _ := repository.GetOrderByCustomer()
	c.JSON(http.StatusOK, customers)
}
