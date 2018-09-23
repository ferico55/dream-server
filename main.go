package main

import (
	"database/sql"
	"fmt"

	"server/config"
	"server/controller"
	localMiddleware "server/middleware"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

var db *sql.DB
var err error

func check(err error) {
	if err != nil {
		fmt.Print("ERROR Occured!!!")
		fmt.Print(err)
	}
}

func main() {
	e := echo.New()

	db, err = sql.Open(config.DriverName, config.ConnectionString)
	defer db.Close()

	rows, err := db.Query("SELECT title FROM dreams")
	check(err)

	var s, data string
	for rows.Next() {
		err = rows.Scan(&data)
		check(err)
		s += data + "\n"
	}

	fmt.Print(s)

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(localMiddleware.ServerHeader)

	// Route => handler
	e.GET("/", controller.GetHandler)

	e.Logger.Fatal(e.Start(":" + config.Port))
}
