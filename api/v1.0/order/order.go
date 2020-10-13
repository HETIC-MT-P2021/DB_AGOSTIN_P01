package order

import (
	"github.com/gin-gonic/gin"
)

func ApplyRoutes(r *gin.RouterGroup) {
	orders := r.Group("/orders")
	{
		//routes for customers
		orders.GET("/", GetAllOrderItems)
	}
}
