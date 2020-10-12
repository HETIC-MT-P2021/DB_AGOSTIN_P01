package endpoint

import (
	"github.com/gin-gonic/gin"
)

func IndexAction(c *gin.Context) {
	c.JSON(200, gin.H{"message": "pong"})
}
