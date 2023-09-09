package handlers

import "gorm.io/gorm"

type UserController struct {
	db *gorm.DB
}

func BuildUserController(db *gorm.DB) *UserController {
	return &UserController{
		db: db,
	}
}
