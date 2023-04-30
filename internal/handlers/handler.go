package handlers

import "github.com/labstack/echo"

type Handler interface {
	Register(router *echo.Router)
}
