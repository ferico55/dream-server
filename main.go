package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"server/config"
	"server/controller"
	localMiddleware "server/middleware"
	"server/model"

	"github.com/dgrijalva/jwt-go"
	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"golang.org/x/crypto/bcrypt"
)

var db *sql.DB
var err error

func check(err error) {
	if err != nil {
		fmt.Print("ERROR Occured!!!")
		fmt.Print(err)
	}
}

func getToken(c echo.Context) error {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": "user",
		"password": "user",
	})

	tokenString, _ := token.SignedString([]byte("secret"))
	fmt.Println(tokenString)
	fmt.Println(token.Signature)
	fmt.Println(token.Valid)

	var err = bcrypt.CompareHashAndPassword([]byte("$2y$10$cM8/.ZdU1sH6DcgYS0m2Iex5I/0o973RkcPEn3BOAs/oOb4PNRGna"), []byte("Fer"))
	if err != nil {
		fmt.Println("INVALID PASSWORD")
	} else {
		fmt.Println("VALID PASSWORD")
	}

	return c.String(200, tokenString)
}

func checkToken(c echo.Context) error {
	authHeader := c.Request().Header.Get("authorization")
	bearerToken := strings.Split(authHeader, " ")

	token, error := jwt.Parse(bearerToken[1], func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("There was an error")
		}
		return []byte("secret"), nil
	})
	if error != nil {
		fmt.Println(error)
		return c.String(400, "ERROR")
	}
	if token.Valid {
		fmt.Println(token.Header)
		fmt.Println(token)
		claims := token.Claims.(jwt.MapClaims)
		// claims := make(jwt.MapClaims)
		fmt.Println(claims["username"])
		// mapClaims := token.Claims(jwt.MapClaims)
		// c := u.(*jwt.Token).Claims.(jwt.MapClaims)
		// claimMap := jwt.MapClaims(token.Claims)
		// token.Claims["username"]
		fmt.Println("TOKEN WAS VALID")
	} else {
		return c.String(401, "Invalid authorization token")
	}
	return c.String(200, "asdf")
}

func AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authHeader := c.Request().Header.Get("authorization")
		bearerToken := strings.Split(authHeader, " ")

		if len(bearerToken) != 2 {
			return c.String(http.StatusUnauthorized, "UNAUTHORIZED")
		}

		if strings.ToLower(bearerToken[0]) != "bearer" {
			return c.String(http.StatusUnauthorized, "UNAUTHORIZED")
		}

		var token, error = jwt.Parse(bearerToken[1], func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("There was an error")
			}
			return []byte(config.Secret), nil
		})

		if error != nil {
			fmt.Println(error)
			return c.String(http.StatusUnauthorized, "UNAUTHORIZED")
		}

		if token.Valid {
			claims := token.Claims.(jwt.MapClaims)
			fmt.Println(claims)
			var id = claims["id"].(string)
			var email = claims["email"].(string)
			var name = claims["name"].(string)
			fmt.Println(id)
			fmt.Println(email)
			fmt.Println(name)
			var idInt, _ = strconv.ParseInt(id, 10, 64)
			var u = model.User{idInt, &name, &email, nil}
			c.Set("user", u)
			return next(c)
		}
		return c.String(http.StatusUnauthorized, "UNAUTHORIZED")
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
	// bisa pake 2 ini buat grouping
	// e.Group("/me", AuthMiddleware)
	// e.Group()
	// e.GET()
	e.GET("/", controller.GetHandler)

	e.POST("/login", controller.LoginHandler)
	e.GET("/me", controller.GetMeHandler, AuthMiddleware)
	// e.GET("/token", getToken)
	// e.GET("/check", checkToken)

	e.Logger.Fatal(e.Start(":" + config.Port))
}
