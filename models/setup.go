package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var DB *gorm.DB

func ConnectDatabase() {
	//This should be accept other connections in the future
	database, err := gorm.Open("sqlite3","localapi.db")

	if err != nil {
		panic("Failed to connect to local db")
	}

	database.AutoMigrate(&Device{})

	DB = database
}
