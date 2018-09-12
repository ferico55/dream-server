package main

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

const port = "8080"

type test struct {
	Ok  string `json:"ok" xml:"ok"`
	Foo int    `json:"foo" xml:"foo"`
}

func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Route => handler
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!\n")
	})
	e.POST("/", func(c echo.Context) error {
		a := test{Ok: "asdf", Foo: 1}
		return c.JSON(http.StatusOK, a)
	})

	e.Logger.Fatal(e.Start(":" + port))
}
