package order

import (
	"github.com/gin-gonic/gin"
	"goevent/database"
	"goevent/database/Models"
	"net/http"
)

func AllOrderItems(c *gin.Context) {
	repository := Models.Repository{Conn: database.DbConn}
	itemsDetails, _ := repository.GetAllItemsInOrder()
	c.JSON(http.StatusOK, itemsDetails)
}
