package main

import (
	"github.com/gocrud/database"
	"github.com/gocrud/routes"
)

func main() {
	//database connecting
	database.InitDB()
	//migration
	database.MigrationTable()
	//seeder

	//running web server
	e := routes.Init()

	// e.Logger.Fatal(e.StartAutoTLS(":443"))
	e.Logger.Fatal(e.Start(":8000"))

}
