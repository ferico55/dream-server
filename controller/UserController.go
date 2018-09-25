package controller

import (
	"encoding/json"
	"net/http"
	"server/config"
	"server/repository"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"golang.org/x/crypto/bcrypt"
)

type requestStruct struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type responseStruct struct {
	Token string `json:"token"`
}

func LoginHandler(c echo.Context) error {
	var decoder = json.NewDecoder(c.Request().Body)
	var r requestStruct
	var err = decoder.Decode(&r)
	if err != nil {
		var e = "Invalid Request Format"
		return responseError(c, http.StatusUnprocessableEntity, &e)
	}

	var user = repository.GetUserByEmail(r.Email)
	if user == nil {
		var e = "Wrong email / password combination"
		return responseError(c, http.StatusUnauthorized, &e)
	}

	err = bcrypt.CompareHashAndPassword([]byte(*user.Password), []byte(r.Password))
	if err != nil {
		var e = "Wrong email / password combination"
		return responseError(c, http.StatusUnauthorized, &e)
	} else {
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"id":    user.ID,
			"email": *user.Email,
			"name":  *user.Name,
		})
		tokenString, _ := token.SignedString([]byte(config.Secret))

		return responseJson(c, http.StatusOK, responseStruct{tokenString})
	}
}
