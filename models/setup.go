package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var DB *gorm.DB

func ConnectDatbase() {
	database, err := gorm.Open("sqlite3", "test.db")

	if err != nil {
		panic("failed to connect to fucking database!")
	}

	database.AutoMigrate(&Plant{})

	DB = database

}
