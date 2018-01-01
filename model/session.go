package model

import (
	"github.com/jinzhu/gorm"
	"github.com/srinathgs/mysqlstore"
)

func GetSessionStore(db *gorm.DB) (*mysqlstore.MySQLStore, error) {
	return mysqlstore.NewMySQLStoreFromConnection(
		db.DB(),
		"sessions",
		"/",
		60*60*24*7,
		[]byte("secret"),
	)
}
