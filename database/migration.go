package database

import (
	"github.com/gocrud/database/models"
)

func MigrationTable() {
	db := InitDB()

	defer db.Close()

	//migrate here
	db.AutoMigrate(&models.UsersTable{})
}
