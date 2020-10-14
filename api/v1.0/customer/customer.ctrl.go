package customer

import (
	"github.com/gin-gonic/gin"
	"goevent/database"
	"goevent/database/models"
	"net/http"
	"strconv"
)

func GetAllCustomers(c *gin.Context) {
	repository := models.Repository{Conn: database.DbConn}
	customers, err := repository.GetAllCustomers()
	if err != nil || len(customers) <= 0 {
		c.JSON(http.StatusNotFound, "Couldn't fetch customers.")
		return
	}
	c.JSON(http.StatusOK, customers)
}

func GetCustomerById(c *gin.Context) {
	repository := models.Repository{Conn: database.DbConn}
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	customer, err := repository.GetCustomer(id)
	if err != nil || customer == nil {
		c.JSON(http.StatusNotFound, "Couldn't fetch customer.")
		return
	}
	c.JSON(http.StatusOK, customer)
}

func GetOrdersByCustomer(c *gin.Context) {
	repository := models.Repository{Conn: database.DbConn}
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	orderSummary, err := repository.GetOrderByCustomer(id)
	if err != nil || orderSummary == nil {
		c.JSON(http.StatusNotFound, "Couldn't fetch orderSummary.")
		return
	}
	c.JSON(http.StatusOK, orderSummary)
}
