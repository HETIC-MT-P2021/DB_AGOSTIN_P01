package api

import (
	"github.com/gin-gonic/gin"
	apiv1 "goevent/api/v1.0"
)

func ApplyRoutes(r *gin.Engine) {
	api := r.Group("/api")
	{
		apiv1.ApplyRoutes(api)
	}
}
