package handlers

import (
	"backend/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func (u *UserController) UserLogout(ctx *gin.Context) {

	//获取参数
	var requestBody struct {
		Username string `json:"username"`
	}
	if err := ctx.ShouldBindJSON(&requestBody); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code":    422,
			"message": "请求数据无效",
		})
		return
	}

	name := requestBody.Username

	//数据验证
	if len(name) == 0 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code":    422,
			"message": "用户名不能为空",
		})
		return
	}

	//判断用户名是否存在
	var user models.User
	u.db.Model(&models.User{}).Where("name = ?", name).First(&user)
	if user.ID == 0 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code":    422,
			"message": "用户不存在",
		})
		return
	}

	result := u.db.Model(&models.User{}).Where("name = ?", name).Update("online", false)
	if result.Error != nil {
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
