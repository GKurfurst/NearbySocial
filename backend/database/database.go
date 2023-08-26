package database

import (
	"backend/models"
	"fmt"
	"github.com/jinzhu/gorm"
)

func InitDB() *gorm.DB {
	driverName := "mysql"
	host := "127.0.0.1"
	port := "3306"
	database := "users_data"
	username := "root"
	password := "zkq200349"
	charset := "utf8"
	args := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=%s&parseTime=true",
		username,
		password,
		host,
		port,
		database,
		charset)

	db, err := gorm.Open(driverName, args)
	if err != nil {
		panic("failed to connect database, err:" + err.Error())
	}

	//迁移
	db.AutoMigrate(&models.User{})

	return db
}
