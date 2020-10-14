package employees

import (
	"github.com/gin-gonic/gin"
	"goevent/database"
	"goevent/database/models"
	"net/http"
	"strconv"
)

func GetEmployees(c *gin.Context) {
	repository := models.Repository{Conn: database.DbConn}
	employees, err := repository.GetAllEmployees()
	if err != nil || len(employees) <= 0 {
		c.JSON(http.StatusNotFound, "Couldn't fetch employees.")
		return
	}
	c.JSON(http.StatusOK, employees)
}

func GetEmployee(c *gin.Context) {
	repository := models.Repository{Conn: database.DbConn}
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	employee, err := repository.GetEmployeeAction(id)
	if err != nil || employee == nil {
		c.JSON(http.StatusNotFound, "Couldn't fetch employee.")
		return
	}
	c.JSON(http.StatusOK, employee)
}
