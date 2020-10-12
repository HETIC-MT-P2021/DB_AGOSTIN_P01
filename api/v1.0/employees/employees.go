package employees

import (
	"github.com/gin-gonic/gin"
)

// ApplyRoutes applies router to the gin Engine
func ApplyRoutes(r *gin.RouterGroup) {
	employees := r.Group("/employees")
	{
		//routes for customers
		employees.GET("/", GetEmployees)
		employees.GET("/:id", GetEmployee)

	}
}
