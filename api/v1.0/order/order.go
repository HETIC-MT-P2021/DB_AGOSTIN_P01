package order

import (
	"github.com/gin-gonic/gin"
)

// ApplyRoutes applies router to the gin Engine
func ApplyRoutes(r *gin.RouterGroup) {
	orders := r.Group("/orders")
	{
		//routes for customers
		orders.GET("/", AllOrderItems)
	}
}
