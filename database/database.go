package database

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gocrud/config"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db  *gorm.DB
	err error
)

func InitDB() *gorm.DB {
	connect := config.GetConfig()

	connectString := connect.DB_USERNAME + ":" + connect.DB_PASSWORD + "@" + "(" + connect.DB_HOST + ":" + connect.DB_PORT + ")/" + connect.DB_NAME + "?charset=utf8&parseTime=True&loc=Local"

	db, err := gorm.Open("mysql", connectString)

	if err != nil {
		fmt.Println(err.Error())
		panic("Failed connect to Database")
	}

	return db
}
