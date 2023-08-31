package handlers

import (
	"backend/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (u *UserController) ApproveFriendRequest(c *gin.Context) {

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

	//判断请求是否存在
	var request models.Request
	if err := u.db.Where("sender_id = ? AND receiver_id = ?", sender.ID, receiver.ID).First(&request).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Request not found",
		})
		return
	}

	//判断该请求是否处于已发送状态
	if request.Status != "pending" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request status",
		})
		return
	}

	//改变请求状态、并添加好友
	request.Status = "approved"
	u.db.Save(&request)

	sender.Friends = append(sender.Friends, receiver)
	receiver.Friends = append(receiver.Friends, sender)

	u.db.Save(&sender)
	u.db.Save(&receiver)

	c.JSON(http.StatusOK, gin.H{
		"message": "Friend request approved",
	})
}
