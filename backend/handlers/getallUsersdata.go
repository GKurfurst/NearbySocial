package handlers

import (
	"backend/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 返回所有已注册用户的信息
func (u *UserController) GetUsers(c *gin.Context) {
	var users []models.User
	u.db.Find(&users)
	c.JSON(http.StatusOK, users)
}
