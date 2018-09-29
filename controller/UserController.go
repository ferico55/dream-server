package controller

import (
	"encoding/json"
	"net/http"
	"server/config"
	"server/model"
	"server/repository"
	"strconv"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"golang.org/x/crypto/bcrypt"
)

type requestStruct struct {
	Name     string `json:"name"`
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
			"id":    strconv.FormatInt(user.ID, 10),
			"email": *user.Email,
			"name":  *user.Name,
		})
		tokenString, _ := token.SignedString([]byte(config.Secret))

		return responseJson(c, http.StatusOK, responseStruct{tokenString})
	}
}

func GetMeHandler(c echo.Context) error {
	user := c.Get("user")
	u, ok := user.(model.User)
	if !ok {
		err := "Invalid Request"
		return responseError(c, http.StatusBadRequest, &err)
	}
	return responseJson(c, http.StatusOK, u)
}

func RegisterHandler(c echo.Context) error {
	var decoder = json.NewDecoder(c.Request().Body)
	var r requestStruct
	var err = decoder.Decode(&r)
	if err != nil {
		var e = "Invalid Request Format"
		return responseError(c, http.StatusUnprocessableEntity, &e)
	}

	var isValid = true
	var e string
	if r.Email == "" {
		isValid = false
		e = "Email Address has to be filled"
	} else if r.Name == "" {
		isValid = false
		e = "Name has to be filled"
	} else if r.Password == "" {
		isValid = false
		e = "Password has to be filled"
	}

	if !isValid {
		return responseError(c, http.StatusUnprocessableEntity, &e)
	}

	checkUser := repository.GetUserByEmail(r.Email)
	if checkUser != nil {
		var e = "This email has been registered, please use another email address."
		return responseError(c, http.StatusUnprocessableEntity, &e)
	}

	password, err := bcrypt.GenerateFromPassword([]byte(r.Password), 10)
	passwordString := string(password[:])
	if err != nil {
		var e = "Unable to encrypt password"
		return responseError(c, http.StatusUnprocessableEntity, &e)
	}

	userID, err := repository.CreateUser(r.Name, r.Email, passwordString)
	if err != nil {
		return responseError(c, http.StatusInternalServerError, nil)
	}

	user := repository.GetUserByID(userID)
	if user == nil {
		return responseError(c, http.StatusInternalServerError, nil)
	}

	return responseJson(c, http.StatusCreated, user)
}
