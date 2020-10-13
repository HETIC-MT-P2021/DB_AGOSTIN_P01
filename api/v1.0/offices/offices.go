package offices

import (
	"github.com/gin-gonic/gin"
)

// ApplyRoutes applies router to the gin Engine
func ApplyRoutes(r *gin.RouterGroup) {
	offices := r.Group("/offices")
	{
		//routes for customers
		offices.GET("/", GetOffices)
		offices.GET("/:id", GetOffice)

	}
}
