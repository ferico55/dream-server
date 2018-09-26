package controller

import (
	"net/http"
	"server/model"
	"server/repository"

	"github.com/labstack/echo"
)

func GetMyDreamHandler(c echo.Context) error {
	user := c.Get("user")
	u, _ := user.(model.User)

	var dreams = repository.GetMyDreams(int(u.ID))
	return responseJson(c, http.StatusOK, dreams)
}
