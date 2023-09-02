package handlers

import (
	"backend/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (u *UserController) GetSendingRequest(c *gin.Context) {

	//获取参数
	userId := c.Param("user_id")

	//判断发送者和接收者是否存在
	var user models.User
	if err := u.db.Where("user_id = ?", userId).First(&user).Error; err != nil {
		c.JSON(400, gin.H{"error": "User not found"})
		return
	}

	//判断是否存在请求
	var requests []models.Request
	if err := u.db.Model(&models.Request{}).Where("status = 'pending' AND receiver_id = ?", user.ID).Find(&requests).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "查询失败",
		})
		return
	}

	if len(requests) == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "No request",
		})
		return
	}

	//返回结果
	c.JSON(http.StatusOK, requests)

}
