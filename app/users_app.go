package app

import (
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gocrud/database"
	"github.com/gocrud/models"
	"github.com/labstack/echo/v4"
)

func AllUsers(c echo.Context) error {
	// database.InitDB()

	db := database.InitDB()

	defer db.Close()

	var res models.Response
	var users []models.UsersTable

	db.Find(&users)
	res.Status = http.StatusOK
	res.Message = "Status OK"
	res.Data = users

	return c.JSON(http.StatusOK, res)
}

func StoreUser(c echo.Context) error {
	db := database.InitDB()

	defer db.Close()

	name := c.Param("name")
	email := c.Param("email")

	db.Create(&models.UsersTable{Name: name, Email: email})

	return c.JSON(http.StatusCreated, map[string]string{"message": "New User is Created!"})
}

func Login(c echo.Context) error {
	db := database.InitDB()

	defer db.Close()

	nameForm := c.FormValue("name")
	emailForm := c.FormValue("email")

	var user models.UsersTable

	db.Where("name = ?", nameForm).Where("email = ?", emailForm).Find(&user)

	if nameForm != user.Name || emailForm != user.Email {
		return echo.ErrUnauthorized
	} else {
		token := jwt.New(jwt.SigningMethodHS256)

		claims := token.Claims.(jwt.MapClaims)
		claims["name"] = user.Name
		claims["email"] = user.Email
		claims["expired"] = time.Now().Add(time.Hour * 72).Unix()

		t, err := token.SignedString([]byte("hai"))

		if err != nil {
			return err
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"token": t,
			"data":  claims,
		})
	}
}
