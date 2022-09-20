package model

type User struct {
	ID       int    `gorm:"primaryKey"`
	Name     string `gorm:"name"`
	Password string `gorm:"password"`
	City     string `gorm:"city"`
}
