package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/richbrowne/go-octo-web/echo-mux/adapters/rest/api/v1/users"
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

func main() {
	e := echo.New()
	log.SetLevel(log.DEBUG)
	if l, ok := e.Logger.(*log.Logger); ok {
		l.SetHeader("${time_rfc3339} ${level} ${prefix} ${long_file} ${line}")
	}

	u := users.NewUsersHandler()
	u.RegisterService(e)

	e.Logger.Fatal(e.Start(":1323"))
}
