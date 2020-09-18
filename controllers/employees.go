package controllers

import (
	"github.com/gin-gonic/gin"
	"goevent/database"
	"goevent/database/Models"
	"net/http"
	"strconv"
)

func GetEmployees(c *gin.Context) {
	repository := Models.Repository{Conn: database.DbConn}
	employees, _ := repository.GetAllEmployees()
	c.JSON(http.StatusOK, employees)
}

func GetEmployee(c *gin.Context) {
	repository := Models.Repository{Conn: database.DbConn}
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	employee, _ := repository.GetEmployeeAction(id)
	c.JSON(http.StatusOK, employee)
}
