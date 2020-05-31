package models

import "github.com/jinzhu/gorm"

type UsersTable struct {
	gorm.Model
	Name  string
	Email string
}
