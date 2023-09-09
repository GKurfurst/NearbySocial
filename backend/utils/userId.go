package utils

import (
	"math/rand"
	"time"
)

func GenerateUserId() int64 {
	rand.Seed(time.Now().UnixNano())
	userID := rand.Int63n(900000) + 100000
	return userID
}
