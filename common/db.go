package common

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB //global variable

func DbConnection(url string) (err error) {
	fmt.Println("Database call")
	DB, err = gorm.Open(postgres.Open(url), &gorm.Config{})

	return
}
