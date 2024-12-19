package db

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"pharmacy/config"

	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() {

	var err error
	if config.Config.DB == "pg" {
		dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
			config.Config.DBHOST, config.Config.DBUSER, config.Config.DBPASSWORD, config.Config.DBNAME, config.Config.DBPORT)
		DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	} else {
		DB, err = gorm.Open(sqlite.Open("phdb.db"), &gorm.Config{})
	}

	if err != nil {
		panic("failed to connect to database: " + err.Error())
	}
}
