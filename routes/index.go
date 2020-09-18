package routes

import (
	"github.com/gin-gonic/gin"
	"goevent/controllers"
)

func addIndexRoutes(rg *gin.RouterGroup) {
	index := rg.Group("/")
	customers := rg.Group("/customers")
	orders := rg.Group("/orders")

	// routes for API endpoint
	index.GET("/", controllers.IndexAction)

	//routes for customers
	customers.GET("/", controllers.AllCustomerAction)
	customers.GET("/:id", controllers.CustomerAction)
	customers.GET("/:id/orders", controllers.OrderByCustomer)

	//routes for order
	orders.GET("/", controllers.AllOrderItems)
}
