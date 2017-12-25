package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql"
)

var db *gorm.DB

func init() {
	//todo specify database connection settings
	dbParam := ""
	db, err := gorm.Open("mysql", dbParam)
	if err != nil {
		panic("failed to connect database")
	}

	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(20)
	db.SingularTable(true)
	db.LogMode(true)
}
