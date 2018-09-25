package controller

import (
	"encoding/json"
	"fmt"

	"github.com/labstack/echo"
)

type requestStruct struct {
	Username string `json: username`
	Password string `json: password`
}

func LoginHandler(c echo.Context) error {
	var decoder = json.NewDecoder(c.Request().Body)
	var t requestStruct
	var err = decoder.Decode(&t)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(t)
	}
	return c.String(200, "")
}
