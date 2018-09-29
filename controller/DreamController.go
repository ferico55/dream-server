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

func GetDreamByIdHandler(c echo.Context) error {
	id := c.Param("id")
	user := c.Get("user")
	u, _ := user.(model.User)

	dream := repository.GetDreamByID(id)
	if dream == nil {
		err := "Dream not found"
		return responseError(c, http.StatusOK, &err)
	}
	if dream.UserID == int(u.ID) {
		return responseJson(c, http.StatusOK, dream)
	}

	err := "This is not your dream"
	return responseError(c, http.StatusForbidden, &err)
}
