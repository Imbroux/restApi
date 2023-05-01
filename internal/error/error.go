package error

import (
	"errors"
	"github.com/labstack/echo"
)

func NewErrorResponse(c echo.Context, errStatus int, message string) error {
	err := errors.New(message)
	_, ok := err.(*echo.HTTPError)
	if !ok {
		report := echo.NewHTTPError(errStatus, err.Error())
		_ = c.JSON(errStatus, report)
	}
	c.Error(errors.New("internal server error"))
	return nil
}
