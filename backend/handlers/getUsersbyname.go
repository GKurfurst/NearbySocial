package handlers

import (
	"backend/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (u *UserController) GetUsersByName(ctx *gin.Context) {

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
	u.db.Preload("Friends").Where("name = ?", name).First(&user)
	if user.ID == 0 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code":    422,
			"message": "用户不存在",
		})
		return
	}

	//返回结果
	ctx.JSON(http.StatusOK, user)

}
