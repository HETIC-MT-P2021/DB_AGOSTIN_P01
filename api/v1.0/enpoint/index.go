package endpoint

import (
	"github.com/gin-gonic/gin"
)

// ApplyRoutes applies router to the gin Engine
func ApplyRoutes(r *gin.RouterGroup) {
	employees := r.Group("/")
	{
		employees.GET("/", IndexAction)
	}
}
