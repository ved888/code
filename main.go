package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"main/common"
	"main/model"
	"main/user"
)

func main() {

	postgresCres := "host=localhost user=postgres password=ved1234 dbname=apiuser port=5432 sslmode=disable"

	err := common.DbConnection(postgresCres)

	if err != nil {
		fmt.Println("error in db connection", err)
		return
	}
	err = common.DB.AutoMigrate(&model.User{})

	if err != nil {
		fmt.Println("error in db connection", err)
		return
	}

	router := gin.Default()

	r := router.Group("/user")

	r.POST("/creat", user.Create)
	r.GET("/get_all", user.GetAll)
	r.GET("/get_single", user.GetSingle)
	r.DELETE("/delete", user.Delete)

	router.Run(":5000")
}
