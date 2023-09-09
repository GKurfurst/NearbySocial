package database

import (
	"backend/models"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	host := "8.130.23.188"
	port := "3306"
	database := "testsocial"
	username := "mysqluser"
	password := ""
	charset := "utf8"
	args := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=%s&parseTime=true",
		username,
		password,
		host,
		port,
		database,
		charset)

	db, err := gorm.Open(mysql.Open(args), &gorm.Config{})
	if err != nil {
		panic("failed to connect database, err:" + err.Error())
	}

	//迁移
	db.AutoMigrate(&models.User{}, &models.Request{})
	return db
}
