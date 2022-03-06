package users

import (
	"net/http"

	"github.com/labstack/gommon/log"

	"github.com/labstack/echo/v4"
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

type UserHandler struct {
	users UserStore
}

func NewUsersHandler() *UserHandler {
	return &UserHandler{
		users: make(UserStore),
	}
}

func (h *UserHandler) addUser(ctx echo.Context) error {
	log.Debug("addUser")
	aUser := new(User)
	if err := ctx.Bind(aUser); err != nil {
		return nil
	}

	sUser := UserDTO{
		Name:    aUser.Name,
		Email:   aUser.Email,
		IsAdmin: false,
	}

	log.Debugf("adding user %+v", sUser)
	if u, exists := h.users[aUser.Name]; !exists {
		h.users[u.Name] = sUser
	}

	return ctx.JSON(http.StatusOK, aUser)
}

func (h *UserHandler) getUser(ctx echo.Context) error {
	log.Debug("getUser")
	name := ctx.Param("name")
	if u, exists := h.users[name]; exists {
		resp := User{
			Name:  u.Name,
			Email: u.Email,
		}
		return ctx.JSON(http.StatusOK, resp)
	} else {
		return echo.ErrNotFound
	}
}

func (h *UserHandler) deleteUser(ctx echo.Context) error {
	name := ctx.Param("name")
	delete(h.users, name)
	return ctx.NoContent(http.StatusNoContent)
}

func (h *UserHandler) RegisterService(e *echo.Echo) {
	e.GET("/users/:id", h.getUser)
	e.POST("/users", h.addUser)
	e.DELETE("/users/:id", h.deleteUser)
}
