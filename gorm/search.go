package gorm

import (
	"gorm.io/gorm"
)

type User struct {
	UserID    int     `json:"user_id" gorm:"primaryKey;column:user_id"`
	SelfIntro string  `json:"self_introduction"`
	Longitude float64 `json:"longtitude"`
	Latitude  float64 `json:"latitude"`
	IsOnline  bool    `json:"is_online"`
	Username  string  `json:"username"`
}

func GetUserByID(db *gorm.DB, id int) (User, error) {
	var user User
	err := db.First(&user, id).Error
	return user, err
}
