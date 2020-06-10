package utils

import (
	"github.com/MVC/database"
	"github.com/MVC/routes"
)

func Init() {
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
