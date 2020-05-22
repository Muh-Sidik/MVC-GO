package models

import "github.com/gocrud/database"

func MigrationTable() {
	db := database.InitDB()

	defer db.Close()

	//migrate here
	db.AutoMigrate(&UsersTable{})
}
