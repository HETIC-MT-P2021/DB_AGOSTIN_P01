package order

import (
	"github.com/gin-gonic/gin"
	"goevent/database"
	"goevent/database/models"
	"net/http"
)

func GetAllOrderItems(c *gin.Context) {
	repository := models.Repository{Conn: database.DbConn}
	itemsDetails, err := repository.GetAllItemsInOrder()
	if err != nil || len(itemsDetails) <= 0 {
		c.JSON(http.StatusNotFound, "Couldn't fetch itemsDetails.")
		return
	}
	c.JSON(http.StatusOK, itemsDetails)
}
