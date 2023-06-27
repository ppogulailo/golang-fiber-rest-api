package database

import (
	"github.com/PogunGun/golang-fiber-rest-api/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	database, err := gorm.Open(postgres.Open("postgresql://postgres:root@localhost:5432/postgres"), &gorm.Config{})

	if err != nil {
		panic("Could not connect to database!")
	}
	DB = database

	database.AutoMigrate(&models.User{})
}
