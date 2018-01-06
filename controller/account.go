package controller

import (
	"github.com/labstack/echo"
	"github.com/letusfly85/echo-example-api/model"
	"net/http"
	"strconv"
)

func SignUp(c echo.Context) (err error) {
	param := new(model.Account)

	if err := c.Bind(param); err != nil {
		return c.JSON(http.StatusBadRequest, "sign up parameter was invalid")
	}

	account := model.Account{}
	if err = account.Create(*param); err != nil {
		return c.JSON(http.StatusInternalServerError, "account already exists")
	}

	SaveSession(c)
	return c.JSON(http.StatusCreated, "account created successfully")
}

func SignIn(c echo.Context) (err error) {
	param := new(model.Account)

	if err := c.Bind(param); err != nil {
		return c.JSON(http.StatusBadRequest, "sign in parameter was invalid")
	}

	account := model.Account{}
	if err = account.Find(*param); err != nil {
		return c.JSON(http.StatusForbidden, err.Error())
	}

	SaveSession(c)
	return c.JSON(http.StatusCreated, "sign in successfully")
}

func SignOut(c echo.Context) error {
	err := DeleteSession(c)

	return err
}

func Retire(c echo.Context) error {
	param := new(model.Account)

	account := model.Account{}
	if err := account.Delete(*param); err != nil {
		return c.JSON(http.StatusInternalServerError, "account not exists")
	}

	return c.JSON(http.StatusCreated, "account deleted successfully")
}

var accountsJson map[string]string
func AccountList(c echo.Context) (err error) {
	accounts := model.Accounts{}
	accounts.List()

	accountsJson = map[string]string {}

	for _, account := range accounts {
		id := strconv.Itoa(int(account.ID))
		accountsJson[id] = account.Email
	}

	return c.JSON(http.StatusOK, accountsJson)
}
