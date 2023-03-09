package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	connection, err := gorm.Open(postgres.Open("postgres://postgres:12345@localhost:5432/assignment3go"), &gorm.Config{})
	if err != nil {
		panic("Could not connect to the db")
	}

	DB = connection
}
