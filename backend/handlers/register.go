package handlers

import (
	"backend/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

func (u *UserController) UserRegister(ctx *gin.Context) {
	//获取参数
	name := ctx.PostForm("name")
	telephone := ctx.PostForm("telephone")
	password := ctx.PostForm("password")

	//数据验证
	if len(name) == 0 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code":    422,
			"message": "用户名不能为空",
		})
		return
	}
	if len(telephone) != 11 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code":    422,
			"message": "手机号必须为11位",
		})
		return
	}
	if len(password) < 6 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code":    422,
			"message": "密码不能少于6位",
		})
		return
	}

	//判断用户名是否存在
	var user models.User
	u.db.Model(&models.User{}).Where("name = ?", name).First(&user)
	if user.ID != 0 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code":    422,
			"message": "用户名已存在",
		})
		return
	}

	//创建用户
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code":    500,
			"message": "密码加密错误",
		})
		return
	}
	newUser := models.User{
		Name:      name,
		Telephone: telephone,
		Password:  string(hashedPassword),
	}
	result := u.db.Model(&models.User{}).Create(&newUser)
	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "注册失败",
		})
		return
	}

	//返回结果
	ctx.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "注册成功",
	})
}
