package user

import (
	"github.com/labstack/echo"
	"net/http"
	error2 "restApi/internal/error"
	"strconv"
)

type Handler struct {
}

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

var users []User

func (h *Handler) Register() *echo.Echo {
	e := echo.New()
	e.GET("/users", getUsers)
	e.GET("/users/:id", getUser)
	e.POST("/users", createUser)
	e.PUT("/users/:id", updateUser)
	e.DELETE("/users/:id", deleteUser)
	return e
}
func getUsers(c echo.Context) error {
	return c.JSON(http.StatusOK, users)
}

func getUser(c echo.Context) error {
	id := c.Param("id")

	for _, user := range users {
		if strconv.Itoa(user.ID) == id {
			return c.JSON(http.StatusOK, user)
		}
	}
	return error2.NewErrorResponse(c, http.StatusNotFound, "user not foun")
}

func createUser(c echo.Context) error {
	user := new(User)
	if err := c.Bind(user); err != nil {
		return error2.NewErrorResponse(c, http.StatusBadRequest, "invalid request body")
		return err
	}
	users = append(users, *user)
	return c.JSON(http.StatusCreated, user)
}

func updateUser(c echo.Context) error {
	id := c.Param("id")

	for i, user := range users {
		if strconv.Itoa(user.ID) == id {
			newUser := new(User)
			if err := c.Bind(newUser); err != nil {
				return err
			}
			users[i] = *newUser
			return c.JSON(http.StatusOK, users[i])
		}
	}
	return error2.NewErrorResponse(c, http.StatusNotFound, "users not found or not update")
}

func deleteUser(c echo.Context) error {
	id := c.Param("id")

	for i, user := range users {
		if strconv.Itoa(user.ID) == id {
			users = append(users[:i], users[i+1:]...)
			return c.NoContent(http.StatusNoContent)
		}
	}
	return error2.NewErrorResponse(c, http.StatusNotFound, "users not found or not Deleted")

}
