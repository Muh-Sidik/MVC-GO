package models

import "github.com/jinzhu/gorm"

type UsersTable struct {
	gorm.Model
	Username string
	Password string
	Status   uint `gorm:"default:0"`
}
