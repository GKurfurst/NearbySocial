package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name      string    `gorm:"varchar(20);not null"`
	UserId    string    `gorm:"not null;unique"`
	Telephone string    `gorm:"varchar(20);not null;unique"`
	Password  string    `gorm:"size:288;not null"`
	Friends   []User    `gorm:"many2many:user_friends;association_jointable_foreignkey:user_id;association_foreignkey:friend_id"`
	Requests  []Request `gorm:"foreignkey:ReceiverId"`
	Online    bool
}

type Request struct {
	gorm.Model
	Sender     User   `gorm:"foreignkey:SenderId"`
	SenderId   string `gorm:"not null"`
	Receiver   User   `gorm:"foreignkey:ReceiverId"`
	ReceiverId string `gorm:"not null"`
	Status     string `gorm:"not null"`
}
