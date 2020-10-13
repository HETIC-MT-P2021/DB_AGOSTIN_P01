package customer

import (
	"github.com/gin-gonic/gin"
)

// ApplyRoutes applies router to the gin Engine
func ApplyRoutes(r *gin.RouterGroup) {
	customers := r.Group("/customer")
	{
		//routes for customers
		customers.GET("/", GetAllCustomers)
		customers.GET("/:id", GetCustomerById)
		customers.GET("/:id/orders", GetOrdersByCustomer)

	}
}
