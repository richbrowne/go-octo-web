package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

// User
type User struct {
	Name  string `json:"name" form:"name" query:"name"`
	Email string `json:"email" form:"email" query:"email"`
}

type UserDTO struct {
	Name    string
	Email   string
	IsAdmin bool
}

func addUser(ctx echo.Context) error {
	u := new(User)
	if err := ctx.Bind(u); err != nil {
		return nil
	}

	user := UserDTO{
		Name:    u.Name,
		Email:   u.Email,
		IsAdmin: false,
	}
	log.Debugf("%+v", user)

	return ctx.JSON(http.StatusOK, u)
}

func getUser(ctx echo.Context) error {
	u := new(User)
	if err := ctx.Bind(u); err != nil {
		return nil
	}

	return ctx.JSON(http.StatusOK, u)
}

func main() {
	e := echo.New()
	e.Debug = true
	log.SetLevel(log.DEBUG)
	if l, ok := e.Logger.(*log.Logger); ok {
		l.SetHeader("${time_rfc3339} ${level} ${prefix} ${long_file} ${line}")
		l.SetLevel(log.DEBUG)
	}
	
	e.GET("/users", getUser)
	e.POST("/users", addUser)

	e.Logger.Fatal(e.Start(":1323"))
}
