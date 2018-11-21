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
	e.POST("/register", controller.RegisterHandler)

	// ME GROUP
	e.GET("me/", controller.GetMeHandler, localMiddleware.AuthMiddleware)
	e.GET("me/dreams", controller.GetMyDreamHandler, localMiddleware.AuthMiddleware)

	// DREAM Group
	e.GET("/dreams/:id", controller.GetDreamByIdHandler, localMiddleware.AuthMiddleware)
	e.POST("/dreams", controller.CreateDreamHandler, localMiddleware.AuthMiddleware)

	// TODO Group
	e.POST("todo/:id/check", controller.CheckTodo, localMiddleware.AuthMiddleware)
	e.POST("todo/:id/uncheck", controller.UncheckTodo, localMiddleware.AuthMiddleware)
	e.POST("/dreams/:id/todo", controller.CreateTodoHandler, localMiddleware.AuthMiddleware)

	e.GET("/error", controller.ErrorHandler)

	e.Logger.Fatal(e.Start(":" + config.Port))
}
