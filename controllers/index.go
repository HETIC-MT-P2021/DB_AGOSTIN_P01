package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func IndexAction(context *gin.Context) {
	context.JSON(http.StatusOK, "index")
}
