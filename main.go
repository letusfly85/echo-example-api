package main

import (
	"github.com/gorilla/sessions"
	"github.com/labstack/echo"
	m "github.com/labstack/echo/middleware"
	"github.com/labstack/echo-contrib/session"

	"github.com/letusfly85/echo-example-api/controller"
	"github.com/letusfly85/echo-example-api/middlewares"
)


func main() {
	e := echo.New()

	e.Use(session.Middleware(sessions.NewCookieStore([]byte("secret"))))
	e.Use(middlewares.CheckSession)
	e.Use(m.RequestID())
	e.Use(m.Logger())

	e.POST("/signUp", controller.SignUp)
	e.POST("/signIn", controller.SignIn)

	e.POST("/signOut", controller.SignOut)
	e.POST("/retire", controller.Retire)

	e.GET("/accounts", controller.AccountList)

	e.Logger.Fatal(e.Start(":3030"))
}
