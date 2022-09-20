package user

import (
	"errors"
	"github.com/gin-gonic/gin"
	"main/common"
	"main/model"
	"net/http"
)

func Create(c *gin.Context) {
	var user model.User
	name := c.Request.FormValue("name")
	password := c.Request.FormValue("password")
	city := c.Request.FormValue("city")

	if name == "" {
		c.AbortWithError(http.StatusBadRequest, errors.New("name is required"))
		return
	}
	if password == "" {
		c.AbortWithError(http.StatusBadRequest, errors.New("password is required"))
		return

	}
	if city == "" {
		c.AbortWithError(http.StatusBadRequest, errors.New("city is required"))
		return
	}

	user.Name = name
	user.Password = password
	user.City = city

	result := common.DB.Create(&user)

	if result.Error != nil {
		c.AbortWithError(http.StatusInternalServerError, result.Error)
		return
	}

	c.JSON(http.StatusOK, user)
}
