package app

import (
	"net/http"

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
