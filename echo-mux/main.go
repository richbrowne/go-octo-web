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

type UserStore map[string]UserDTO

var users UserStore

func addUser(ctx echo.Context) error {
	aUser := new(User)
	if err := ctx.Bind(aUser); err != nil {
		return nil
	}

	sUser := UserDTO{
		Name:    aUser.Name,
		Email:   aUser.Email,
		IsAdmin: false,
	}

	if u, exists := users[aUser.Name]; !exists {
		users[u.Name] = sUser
	}

	return ctx.JSON(http.StatusOK, aUser)
}

func getUser(ctx echo.Context) error {
	name := ctx.Param("name")
	if u, exists := users[name]; exists {
		resp := User{
			Name:  u.Name,
			Email: u.Email,
		}
		return ctx.JSON(http.StatusOK, resp)
	} else {
		return echo.ErrNotFound
	}
}

func deleteUser(ctx echo.Context) error {
	name := ctx.Param("name")
	delete(users, name)
	return ctx.NoContent(http.StatusNoContent)
}

func main() {
	e := echo.New()
	log.SetLevel(log.DEBUG)
	if l, ok := e.Logger.(*log.Logger); ok {
		l.SetHeader("${time_rfc3339} ${level} ${prefix} ${long_file} ${line}")
	}

	users = make(UserStore)
	e.GET("/users/:id", getUser)
	e.POST("/users", addUser)
	e.DELETE("/users/:id", deleteUser)

	e.Logger.Fatal(e.Start(":1323"))
}
