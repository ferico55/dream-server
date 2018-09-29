package controller

import (
	"fmt"

	"github.com/labstack/echo"
)

func testPackageFunc() {
	fmt.Print("test")
}

type responseWrapper struct {
	Success bool        `json:"success"`
	Error   *string     `json:"error"`
	Data    interface{} `json:"data"`
}

func responseJson(context echo.Context, statusCode int, data interface{}) error {
	var response = responseWrapper{true, nil, data}
	return context.JSON(statusCode, response)
}

func responseError(context echo.Context, statusCode int, e *string) error {
	if e == nil {
		err := "Something went wrong with our server, please try again later."
		e = &err
	}
	var response = responseWrapper{false, e, nil}
	return context.JSON(statusCode, response)
}
