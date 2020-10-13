package errors

import "github.com/gin-gonic/gin"

func HandlerError() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		errorToPrint := c.Errors.ByType(gin.ErrorTypePublic).Last()
		if errorToPrint != nil && errorToPrint.Meta != nil {
			c.JSON(500, errorToPrint.Meta)
		}
	}
}
