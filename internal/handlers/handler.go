package handlers

import "github.com/labstack/echo"

type HandlerInterface interface {
	Register() *echo.Echo
}
