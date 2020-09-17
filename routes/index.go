package routes

import (
	"github.com/gin-gonic/gin"
	"goevent/controllers"
)

func addIndexRoutes(rg *gin.RouterGroup) {
	index := rg.Group("/")
	customers := rg.Group("/customers")

	// routes for API endpoint
	index.GET("/", controllers.IndexAction)

	//routes for customers
	customers.GET("/", controllers.CustomerDetailsAction)

}
