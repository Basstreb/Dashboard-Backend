package database

import (
	"dashboard/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	connection, err := gorm.Open(mysql.Open("root:root@tcp(localhost:3306)/dash-dev?charset=utf8mb4&parseTime=True"), &gorm.Config{})

	if err != nil {
		panic("Could nor connect to the database")
	}

	DB = connection

	connection.AutoMigrate(&models.UserData{})
	connection.AutoMigrate(&models.ClientData{})
	connection.AutoMigrate(&models.OffersData{})
	connection.AutoMigrate(&models.CommonCosts{})
	connection.AutoMigrate(&models.StaffCosts{})
	connection.AutoMigrate(&models.TypoCosts{})
}
