package database

import (
	"fmt"

	// _ "github.com/lib/pq"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gocrud/config"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	// _ "github.com/jinzhu/gorm/dialects/postgres"
)

var (
	db  *gorm.DB
	err error
)

func InitDB() *gorm.DB {
	connect := config.GetConfig()

	connectStringMysql := connect.DB_USERNAME + ":" + connect.DB_PASSWORD + "@" + "(" + connect.DB_HOST + ":" + connect.DB_PORT + ")/" + connect.DB_NAME + "?charset=utf8&parseTime=True&loc=Local"
	// connectStringPostgresql := "host=" + connect.DB_HOST + "port=" + connect.DB_PORT + "user=" + connect.DB_USERNAME + "dbname=" + connect.DB_NAME + "password=" + connect.DB_PASSWORD

	//Mysql
	db, err := gorm.Open("mysql", connectStringMysql)
	//Postgresql
	// db, err := gorm.Open("postgres", connectStringPostgresql)

	if err != nil {
		fmt.Println(err.Error())
		panic("Failed connect to Database")
	}

	return db
}
