package database

import (
	"github.com/MVC/database/models"
)

func MigrationTable() {
	db := InitDB()

	defer db.Close()

	//migrate here
	db.AutoMigrate(&models.UsersTable{})
}
