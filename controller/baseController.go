package controller

import (
	"fmt"

	"github.com/labstack/echo"
)

func testPackageFunc() {
	fmt.Print("test")
}

type responseWrapper struct {
	success bool        `json: success`
	error   []string    `json: error`
	data    interface{} `json: data`
}

func responseJson(context echo.Context, statusCode int, data interface{}) {
	// TODO: additional logging or handling here
	var response = responseWrapper{true, nil, data}
	context.JSON(statusCode, response)
}

func responseError(context echo.Context, statusCode int, errors []string) {
	var response = responseWrapper{false, errors, nil}
	context.JSON(statusCode, response)
}
