package models

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Name      string `gorm:"varchar(20);not null"`
	UserId    string `gorm:"size:288;not null;"`
	Telephone string `gorm:"varchar(20);not null;unique"`
	Password  string `gorm:"size:288;not null"`
}
