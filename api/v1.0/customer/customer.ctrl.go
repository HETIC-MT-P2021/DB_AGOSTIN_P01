package customer

import (
	"github.com/gin-gonic/gin"
	"goevent/database"
	"goevent/database/Models"
	"net/http"
	"strconv"
)

func GetAllCustomers(c *gin.Context) {
	repository := Models.Repository{Conn: database.DbConn}
	customers, _ := repository.GetAllCustomers()
	c.JSON(http.StatusOK, customers)
}

func GetCustomerById(c *gin.Context) {
	repository := Models.Repository{Conn: database.DbConn}
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	customer, _ := repository.GetCustomer(id)
	c.JSON(http.StatusOK, customer)
}

func GetOrdersByCustomer(c *gin.Context) {
	repository := Models.Repository{Conn: database.DbConn}
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	orderSummary, _ := repository.GetOrderByCustomer(id)
	c.JSON(http.StatusOK, orderSummary)
}
