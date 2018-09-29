package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"server/model"
	"server/repository"

	"github.com/labstack/echo"
)

type todoRequestStruct struct {
	Title string `json:"title"`
}

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
		return responseError(c, http.StatusInternalServerError, nil)
	}

	return responseJson(c, http.StatusOK, nil)
}

func UncheckTodo(c echo.Context) error {
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

	err := repository.UncheckTodo(id)
	if err != nil {
		fmt.Println(err)
		return responseError(c, http.StatusInternalServerError, nil)
	}

	return responseJson(c, http.StatusOK, nil)
}

func CreateTodoHandler(c echo.Context) error {
	id := c.Param("id")

	user := c.Get("user")
	u, _ := user.(model.User)

	var decoder = json.NewDecoder(c.Request().Body)
	var r todoRequestStruct
	err := decoder.Decode(&r)
	if err != nil {
		var e = "Invalid Request Format"
		return responseError(c, http.StatusUnprocessableEntity, &e)
	}

	ownerID := repository.GetDreamOwnerID(id)
	if int64(ownerID) != u.ID {
		e := "This is not your dream!"
		return responseError(c, http.StatusForbidden, &e)
	}

	todoID, err := repository.CreateTodo(r.Title, id)
	if err != nil {
		return responseError(c, http.StatusInternalServerError, nil)
	}

	todo := repository.GetTodoByID(todoID)
	return responseJson(c, http.StatusCreated, todo)
}
