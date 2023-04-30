package main

import (
	"github.com/labstack/echo"
	"net/http"
	"strconv"
)

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

var users []User

func main() {
	e := echo.New()

	e.GET("/users", func(c echo.Context) error {
		return c.JSON(http.StatusOK, users)
	})

	e.POST("/users", func(c echo.Context) error {
		user := new(User)
		if err := c.Bind(user); err != nil {
			return err
		}
		users = append(users, *user)
		return c.JSON(http.StatusCreated, user)
	})
	e.DELETE("/users/:id", func(c echo.Context) error {
		id := c.Param("id")
		for i, user := range users {
			if strconv.Itoa(user.ID) == id {
				users = append(users[:i], users[i+1:]...)
				return c.NoContent(http.StatusNoContent)
			}
		}
		return c.JSON(http.StatusNotFound, map[string]string{"error": "user not found"})
	})
	e.PUT("/users/:id", func(c echo.Context) error {

		return c.JSON(http.StatusNotFound, map[string]string{"error": "user not found"})
	})

	e.Logger.Fatal(e.Start(":8080"))
}
