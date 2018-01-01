package middlewares

import (
	"net/http"
	"github.com/labstack/echo"
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

/*
func SetSession(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
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
		session.Save(request, writer)

		return next(c)
	}
}
*/

func CheckSession(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		request := c.Request()
		if c.Request().RequestURI == "/signIn" || c.Request().RequestURI == "/signUp" {
			return next(c)
		}

		session, err := store.Get(request, "Authorize")
		if err != nil {
			return echo.NewHTTPError(
				http.StatusUnauthorized,
				nil,
			)
		}
		authorized, ok := session.Values["authorized"].(bool)
		if !ok || authorized == false {
			return echo.NewHTTPError(
				http.StatusUnauthorized,
				nil,
			)
		}

		return next(c)
	}
}