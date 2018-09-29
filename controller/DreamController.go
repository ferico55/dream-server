package controller

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"server/model"
	"server/repository"
	"strconv"

	"github.com/labstack/echo"
)

type dreamRequestStruct struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	ImageURI    string `json:"image_uri"`
}

func GetMyDreamHandler(c echo.Context) error {
	user := c.Get("user")
	u, _ := user.(model.User)

	var dreams = repository.GetMyDreams(int(u.ID))
	return responseJson(c, http.StatusOK, dreams)
}

func GetDreamByIdHandler(c echo.Context) error {
	id := c.Param("id")
	user := c.Get("user")
	u, _ := user.(model.User)

	dream := repository.GetDreamByID(id)
	if dream == nil {
		err := "Dream not found"
		return responseError(c, http.StatusOK, &err)
	}
	if dream.UserID == int(u.ID) {
		return responseJson(c, http.StatusOK, dream)
	}

	err := "This is not your dream"
	return responseError(c, http.StatusForbidden, &err)
}

func CreateDreamHandler(c echo.Context) error {
	user := c.Get("user")
	u, _ := user.(model.User)

	var decoder = json.NewDecoder(c.Request().Body)
	var r dreamRequestStruct
	var err = decoder.Decode(&r)
	if err != nil {
		var e = "Invalid Request Format"
		return responseError(c, http.StatusUnprocessableEntity, &e)
	}

	if r.Title == "" {
		var e = "Title field is required"
		return responseError(c, http.StatusUnprocessableEntity, &e)
	}

	imageURI := r.ImageURI
	if imageURI == "" {
		imageURI = getRandomImage()
	}

	dreamID, err := repository.CreateDream(r.Title, r.Description, imageURI, u.ID)
	if err != nil {
		return responseError(c, http.StatusInternalServerError, nil)
	}
	dream := repository.GetDreamByID(strconv.FormatInt(dreamID, 10))

	return responseJson(c, http.StatusCreated, dream)
}

func getRandomImage() string {
	arr := [6]string{
		"https://di5fgdew4nptq.cloudfront.net/api2/media/images/ed43201d-7bc8-e711-80d4-a0369fdf7ce4?width=1200",
		"https://i0.wp.com/alycevayleauthor.com/wp-content/uploads/2016/10/Star-Dream-Meaning-The-Meaning-of-Stars-in-Your-Dream_stars-in-the-sky.jpg?resize=600%2C350",
		"https://itsdanielslife.files.wordpress.com/2016/05/c__data_users_defapps_appdata_internetexplorer_temp_saved-images_dreams.jpg?w=600",
		"http://dreamon.world/img/events/01.jpg",
		"https://www.goalcast.com/wp-content/uploads/2016/05/dream-quote1.jpg",
		"https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcStYgRXz45sN_vJ1j07-ppkFthacxUGorE4i1DBI5n0hvC8otq",
	}

	index := rand.Int() % 6

	return arr[index]
}
