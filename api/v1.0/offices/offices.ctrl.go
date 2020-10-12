package offices

import (
	"github.com/gin-gonic/gin"
	"goevent/database"
	"goevent/database/Models"
	"net/http"
	"strconv"
)

func GetOfficesAction(c *gin.Context) {
	repository := Models.Repository{Conn: database.DbConn}
	offices, _ := repository.GetOfficesAction()
	c.JSON(http.StatusOK, offices)
}

func GetOfficeAction(c *gin.Context) {
	repository := Models.Repository{Conn: database.DbConn}
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	offices, _ := repository.GetOfficeAction(id)
	c.JSON(http.StatusOK, offices)
}
