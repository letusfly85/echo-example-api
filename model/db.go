package model

import (
	"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql"
)

var db *gorm.DB

func init() {
	//todo specify database connection settings
	initDb, err := gorm.Open("mysql", "root:password@/example?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database")
	}
	db = initDb

	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(20)
	db.SingularTable(true)
	db.LogMode(true)

	db.AutoMigrate(&Account{})
}

func GetDB() *gorm.DB {
	return db
}
