package controller

import (
	"net/http"

	"github.com/labstack/echo"
)

// ErrorHandler dummy handler to create error
func ErrorHandler(c echo.Context) error {
	error := "shit happens!"
	return responseError(c, http.StatusInternalServerError, &error)
}
