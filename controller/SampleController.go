package controller

import (
	"net/http"
	"server/repository"

	"github.com/labstack/echo"
)

// GetHandler default template handler example
func GetHandler(c echo.Context) error {

	var dreams = repository.GetAllDreams()
	return c.JSON(http.StatusOK, dreams)
}
