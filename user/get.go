package user

import (
	"errors"
	"github.com/gin-gonic/gin"
	common "main/common"
	"main/model"
	"net/http"
)

func GetAll(c *gin.Context) {
	var user []model.User
	result := common.DB.Model(model.User{}).Find(&user)

	if result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}
	c.JSON(http.StatusOK, user)
}
func GetSingle(c *gin.Context) {

	var customer model.User
	id := c.Request.FormValue("id")

	if id == "" {

		c.AbortWithError(http.StatusBadRequest, errors.New("id is required"))
		return
	}
	result := common.DB.Model(model.User{}).Where("id", id).First(&customer)

	if result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	c.JSON(http.StatusOK, customer)

}
