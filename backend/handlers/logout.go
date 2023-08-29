package handlers

import (
	"backend/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func (u *UserController) UserLogout(ctx *gin.Context) {
	var user models.User
	user.Online = false
	if err := u.db.Model(&models.User{}).Save(&user).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "更新用户在线状态失败",
		})
		return
	}

	//返回结果
	ctx.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "用户状态已更新",
	})

	// 通过设置Expires为过去的时间来立即使Token失效
	expiredTime := time.Now().Add(-time.Hour) // 设置为当前时间之前的一个小时

	// 生成 Token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp": expiredTime.Unix(),
		// 其他自定义的 Token 数据
	})

	// 签名 Token
	tokenString, err := token.SignedString([]byte("your_secret_key"))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "生成 Token 失败",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "登出成功",
		"token":   tokenString,
	})
}
