package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() {
	_, err := gorm.Open(postgres.Open("postgresql://postgres:root@localhost:5432/postgres"), &gorm.Config{})

	if err != nil {
		panic("Could not connect to database!")
	}
}
