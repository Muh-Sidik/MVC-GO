package main

import (
	"github.com/gocrud/database"
	"github.com/gocrud/models"
	"github.com/gocrud/routes"
)

func main() {
	//database connecting
	database.InitDB()
	//migration
	models.MigrationTable()

	//running web server
	e := routes.Init()

	e.Logger.Fatal(e.Start(":8000"))

}
