package handlers

import (
	"backend/models"
	"backend/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (u *UserController) RemoveFriend(c *gin.Context) {

	//获取参数
	userID := c.PostForm("user_id")
	friendID := c.PostForm("friend_id")

	//判断发送者和接收者是否存在
	var user models.User
	if err := u.db.Preload("Friends").Where("user_id = ?", userID).First(&user).Error; err != nil {
		c.JSON(400, gin.H{"error": "User not found"})
		return
	}

	var friend models.User
	if err := u.db.Preload("Friends").Where("user_id = ?", friendID).First(&friend).Error; err != nil {
		c.JSON(400, gin.H{"error": "Friend not found"})
		return
	}

	if !utils.IsFriend(user, friend) {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Not a friend",
		})
		return
	}

	// 删除好友关系
	if err := u.db.Model(&user).Association("Friends").Delete(&friend); err != nil {
		c.JSON(400, gin.H{"error": "Failed to delete Friend from user"})
		return
	}

	// 删除好友关系（反向）
	if err := u.db.Model(&friend).Association("Friends").Delete(&user); err != nil {
		c.JSON(400, gin.H{"error": "Failed to delete Friend from friend"})
		return
	}

	u.db.Save(&user)
	u.db.Save(&friend)

	c.JSON(http.StatusOK, gin.H{
		"message": "Friend removed",
	})
}
