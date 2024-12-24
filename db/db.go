package db

import (
	"fmt"
	"pharmacy/config"

	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/sqlite"
	_ "github.com/lib/pq"
)

var DB *gorm.DB

func Init() {

	var err error
	if config.Config.DB == "to be changed" {
		dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
			config.Config.DBHOST, config.Config.DBUSER, config.Config.DBPASSWORD, config.Config.DBNAME, config.Config.DBPORT)
		DB, err = gorm.Open("postgres", dsn)
	} else {
		DB, err = gorm.Open("sqlite3", "phdb")
	}

	if err != nil {
		panic("failed to connect to database")
	}
}
