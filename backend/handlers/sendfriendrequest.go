package handlers

import (
	"backend/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (u *UserController) SendFriendRequest(c *gin.Context) {

	//获取参数
	var requestBody struct {
		SenderId   string `json:"sender_id"`
		ReceiverId string `json:"receiver_id"`
	}
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"code":    422,
			"message": "请求数据无效",
		})
		return
	}

	senderID := requestBody.SenderId
	receiverID := requestBody.ReceiverId

	//判断发送者和接收者是否存在
	var sender models.User
	if err := u.db.Where("user_id = ?", senderID).First(&sender).Error; err != nil {
		c.JSON(400, gin.H{"error": "Sender not found"})
		return
	}

	var receiver models.User
	if err := u.db.Where("user_id = ?", receiverID).First(&receiver).Error; err != nil {
		c.JSON(400, gin.H{"error": "Receiver not found"})
		return
	}

	//判断该请求是否已存在
	var existingRequest models.Request
	if err := u.db.Where("sender_id = ? AND receiver_id = ?", sender.ID, receiver.ID).First(&existingRequest).Error; err == nil {
		c.JSON(400, gin.H{"error": "Friend request already exists"})
		return
	}

	// 创建好友请求
	request := models.Request{
		Sender:     sender,
		SenderId:   strconv.Itoa(int(sender.ID)),
		Receiver:   receiver,
		ReceiverId: strconv.Itoa(int(receiver.ID)),
		Status:     "pending",
	}
	if err := u.db.Create(&request).Error; err != nil {
		c.JSON(500, gin.H{"error": "Failed to add friend"})
		return
	}

	c.JSON(200, gin.H{"message": "Friend request sent"})
}
