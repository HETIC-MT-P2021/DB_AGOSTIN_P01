package customer

import (
	"github.com/gin-gonic/gin"
)

// ApplyRoutes applies router to the gin Engine
func ApplyRoutes(r *gin.RouterGroup) {
	customers := r.Group("/customer")
	{
		//routes for customers
		customers.GET("/", AllCustomersAction)
		customers.GET("/:id", FindByIdCustomerAction)
		customers.GET("/:id/orders", OrderByCustomer)

	}
}
