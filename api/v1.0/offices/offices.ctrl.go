package offices

import (
	"github.com/gin-gonic/gin"
	"goevent/database"
	"goevent/database/models"
	"net/http"
	"strconv"
)

func GetOffices(c *gin.Context) {
	repository := models.Repository{Conn: database.DbConn}
	offices, err := repository.GetOfficesAction()
	if err != nil || len(offices) <= 0 {
		c.JSON(http.StatusNotFound, "Couldn't fetch offices.")
		return
	}
	c.JSON(http.StatusOK, offices)
}

func GetOffice(c *gin.Context) {
	repository := models.Repository{Conn: database.DbConn}
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	office, err := repository.GetOfficeAction(id)
	if err != nil || office == nil {
		c.JSON(http.StatusNotFound, "Couldn't fetch office.")
		return
	}
	c.JSON(http.StatusOK, office)
}
