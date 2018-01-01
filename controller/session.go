package controller

import (
	"github.com/labstack/echo"
	"net/http"
	"github.com/srinathgs/mysqlstore"
	"github.com/letusfly85/echo-example-api/model"
)

var store *mysqlstore.MySQLStore

func init() {
	db := model.GetDB()
	dbStore, err := model.GetSessionStore(db)
	if err != nil {
		panic(err.Error())
	}

	store = dbStore
}

func SaveSession(c echo.Context) (err error) {
	request := c.Request()
	writer := c.Response().Writer

	session, err := store.Get(request, "Authorize")
	if err != nil {
		return echo.NewHTTPError(
			http.StatusUnauthorized,
			nil,
		)
	}
	session.Values["authorized"] = true
	err = session.Save(request, writer)

	return err
}
