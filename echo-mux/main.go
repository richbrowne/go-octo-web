package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/richbrowne/go-octo-web/echo-mux/adapters/rest/api/v1/users"
)

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
