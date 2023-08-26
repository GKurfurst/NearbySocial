package handlers

import "github.com/jinzhu/gorm"

type UserController struct {
	db *gorm.DB
}

func BuildUserController(db *gorm.DB) *UserController {
	return &UserController{
		db: db,
	}
}
