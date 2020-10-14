package customer

import (
	"github.com/gin-gonic/gin"
	"goevent/database"
	"goevent/database/models"
	"log"
	"net/http"
	"strconv"
)

func GetAllCustomers(c *gin.Context) {
	repository := models.Repository{Conn: database.DbConn}
	customers, err := repository.GetAllCustomers()

	log.Println(err)
	log.Println(customers)

	c.JSON(http.StatusOK, customers)
}

func GetCustomerById(c *gin.Context) {
	repository := models.Repository{Conn: database.DbConn}
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	customer, _ := repository.GetCustomer(id)
	c.JSON(http.StatusOK, customer)
}

func GetOrdersByCustomer(c *gin.Context) {
	repository := models.Repository{Conn: database.DbConn}
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	orderSummary, _ := repository.GetOrderByCustomer(id)
	c.JSON(http.StatusOK, orderSummary)
}
