package apiv1

import (
	"github.com/gin-gonic/gin"
	"goevent/api/v1.0/customer"
	"goevent/api/v1.0/employees"
	endpoint "goevent/api/v1.0/enpoint"
	"goevent/api/v1.0/offices"
)

func ApplyRoutes(r *gin.RouterGroup) {
	v1 := r.Group("/v1")
	{
		customer.ApplyRoutes(v1)
		employees.ApplyRoutes(v1)
		endpoint.ApplyRoutes(v1)
		offices.ApplyRoutes(v1)
	}
}
