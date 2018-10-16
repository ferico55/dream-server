package middleware

import (
	"fmt"
	"net/http"
	"server/config"
	"server/model"
	"strconv"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

// ASD
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
			var id = claims["id"].(string)
			var email = claims["email"].(string)
			var name = claims["name"].(string)
			var idInt, _ = strconv.ParseInt(id, 10, 64)
			var u = model.User{idInt, &name, &email, nil}
			c.Set("user", u)
			return next(c)
		}
		return c.String(http.StatusUnauthorized, "UNAUTHORIZED")
	}
}
