package order

import (
	"github.com/gin-gonic/gin"
	"goevent/database"
	"goevent/database/models"
	"net/http"
)

func GetAllOrderItems(c *gin.Context) {
	repository := models.Repository{Conn: database.DbConn}
	itemsDetails, _ := repository.GetAllItemsInOrder()
	c.JSON(http.StatusOK, itemsDetails)
}
