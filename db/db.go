package db

import (
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/sqlite"
)

var DB *gorm.DB

func Init() {
    var err error
    DB, err = gorm.Open("sqlite3", "phdb.db")
    if err != nil {
        panic("failed to connect database")
    }
}
