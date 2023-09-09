package handlers

import (
	"backend/models"
	"backend/utils"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"net/http"
)

func (u *UserController) HandleChatHistory(c *gin.Context) {

	//获取参数
	userId := c.Param("userId")
	friendId := c.Param("friendId")

	//判断发送者和接收者是否存在
	var user models.User
	if err := u.db.Preload("Friends").Where("user_id = ?", userId).First(&user).Error; err != nil {
		c.JSON(400, gin.H{"error": "User not found"})
		return
	}

	var friend models.User
	if err := u.db.Preload("Friends").Where("user_id = ?", friendId).First(&friend).Error; err != nil {
		c.JSON(400, gin.H{"error": "Friend not found"})
		return
	}

	if !utils.IsFriend(user, friend) {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Not a friend",
		})
		return
	}

	redisClient := redis.NewClient(&redis.Options{
		Addr:     "8.130.23.188:6379",
		Password: "Bjlgdxredis",
		DB:       0,
	})

	// 从Redis列表中获取历史消息
	result, err := redisClient.LRange(c, userId+":"+friendId+"_history", 0, -1).Result()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to fetch history",
		})
		return
	}

	var messages []models.Message
	for _, msg := range result {
		var message models.Message
		json.Unmarshal([]byte(msg), &message)
		messages = append(messages, message)
	}

	c.JSON(http.StatusOK, gin.H{
		"history": messages,
	})
}
