package user

import (
	"errors"
	"github.com/gin-gonic/gin"
	"main/common"
	"main/model"
	"net/http"
)

func Delete(c *gin.Context) {

	var user model.User

	id := c.Request.FormValue("id")
	if id == "" {

		c.AbortWithError(http.StatusBadRequest, errors.New("id is required"))
		return
	}
	result1 := common.DB.Model(model.User{}).Where("id", id).First(&user)

	if result1.Error != nil {
		c.AbortWithError(http.StatusNotFound, result1.Error)
		return
	}

	result := common.DB.Unscoped().Delete(&user) //hard delete

	if result.Error != nil {
		c.AbortWithError(http.StatusInternalServerError, result.Error)
		return
	}

	c.JSON(http.StatusOK, user)
}
