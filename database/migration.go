package database

import (
	"github.com/gocrud/database"
	"github.com/gocrud/database/models"
)

func MigrationTable() {
	db := database.InitDB()

	defer db.Close()

	//migrate here
	db.AutoMigrate(&models.UsersTable{})
}
