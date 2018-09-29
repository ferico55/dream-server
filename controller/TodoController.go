package controller

import (
	"fmt"
	"net/http"
	"server/model"
	"server/repository"

	"github.com/labstack/echo"
)

func CheckTodo(c echo.Context) error {
	id := c.Param("id")
	user := c.Get("user")
	u, _ := user.(model.User)

	ownerID := repository.GetTodoOwnerID(id)
	if ownerID == 0 {
		err := "Todo item does not exists"
		return responseError(c, http.StatusUnprocessableEntity, &err)
	} else if ownerID != u.ID {
		err := "You are not the owner of this Todo item"
		return responseError(c, http.StatusForbidden, &err)
	}

	err := repository.CheckTodo(id)
	if err != nil {
		fmt.Println(err)
		err := "Something went wrong"
		return responseError(c, http.StatusInternalServerError, &err)
	}

	return responseJson(c, http.StatusOK, nil)
}
