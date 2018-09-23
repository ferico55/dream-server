package controller

import (
	"net/http"

	"github.com/labstack/echo"
)

// GetHandler default template handler example
func GetHandler(c echo.Context) error {
	testPackageFunc()
	return c.String(http.StatusOK, "Hello, World!\n")
}
