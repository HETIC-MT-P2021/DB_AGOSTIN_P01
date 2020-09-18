package controllers

import (
	"github.com/gin-gonic/gin"
	"goevent/database"
	"goevent/database/Models"
	"net/http"
	"strconv"
)

func AllCustomerAction(c *gin.Context) {
	repository := Models.Repository{Conn: database.DbConn}
	customers, _ := repository.GetAllCustomers()
	c.JSON(http.StatusOK, customers)
}

func CustomerAction(c *gin.Context) {
	repository := Models.Repository{Conn: database.DbConn}
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	customer, _ := repository.GetCustomer(id)
	c.JSON(http.StatusOK, customer)
}

func OrderByCustomer(c *gin.Context) {
	repository := Models.Repository{Conn: database.DbConn}
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	orderSummary, _ := repository.GetOrderByCustomer(id)
	c.JSON(http.StatusOK, orderSummary)
}
