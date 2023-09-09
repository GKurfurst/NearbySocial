package utils

import "backend/models"

func IsFriend(user models.User, friend models.User) bool {
	for _, u := range user.Friends {
		if u.ID == friend.ID {
			return true
		}
	}
	return false
}
