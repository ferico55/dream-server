package main

import (
	"server/config"
	"server/controller"
	localMiddleware "server/middleware"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(localMiddleware.ServerHeader)

	// Route => handler
	// bisa pake 2 ini buat grouping
	// e.Group("/me", AuthMiddleware)
	// e.Group()==

	e.POST("/login", controller.LoginHandler)
	e.GET("/me", controller.GetMeHandler, localMiddleware.AuthMiddleware)

	e.Logger.Fatal(e.Start(":" + config.Port))
}
