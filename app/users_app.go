package app

import (
	"fmt"
	"net/http"

	"github.com/MVC/database"
	"github.com/MVC/database/models"
	"github.com/MVC/database/models/response"
	"github.com/MVC/utils/auth"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

var db = database.InitDB()

type ErrorResponse response.ErrorResponse

func AllUsers(c echo.Context) error {
	var res response.Response
	var users []models.UsersTable

	db.Find(&users)
	res.Status = http.StatusOK
	res.Message = "Status OK"
	res.Data = users

	return c.JSON(http.StatusOK, res)
}

func Register(c echo.Context) error {
	var res response.Response

	username := c.FormValue("username")
	password := c.FormValue("password")

	pass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		fmt.Println(err)
		err := ErrorResponse{
			Info: "Password Encryption Failed!",
		}

		return c.JSON(http.StatusBadRequest, err)
	}

	createdUser := db.Create(&models.UsersTable{
		Username: username,
		Password: string(pass),
	})

	var errMessage = createdUser.Error

	if errMessage != nil {
		fmt.Println(errMessage)
	}

	res.Status = http.StatusCreated
	res.Message = "Register Success!"
	res.Data = createdUser

	return c.JSON(http.StatusCreated, res)

}

func Login(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("passsword")

	res := auth.FindUser(username, password)

	return c.JSON(http.StatusOK, res)
}

// func Login2(c echo.Context) error {
// 	defer db.Close()

// 	nameForm := c.FormValue("username")
// 	passForm := c.FormValue("password")

// 	var user models.UsersTable

// 	db.Where("username = ?", nameForm).Where("pass = ?", emailForm).Find(&user)

// 	if nameForm != user.Name || emailForm != user.Email {
// 		return echo.ErrUnauthorized
// 	} else {
// 		token := jwt.New(jwt.SigningMethodHS256)

// 		claims := token.Claims.(jwt.MapClaims)
// 		claims["name"] = user.Name
// 		claims["email"] = user.Email
// 		claims["expired"] = time.Now().Add(time.Hour * 72).Unix()

// 		t, err := token.SignedString([]byte("hai"))

// 		if err != nil {
// 			return err
// 		}

// 		return c.JSON(http.StatusOK, map[string]interface{}{
// 			"token": t,
// 			"data":  claims,
// 		})
// 	}
// }
