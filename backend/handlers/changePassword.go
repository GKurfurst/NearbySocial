package handlers

import (
	"backend/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

func (u *UserController) UserPasswordChange(ctx *gin.Context) {

	//获取参数
	name := ctx.PostForm("name")
	telephone := ctx.PostForm("telephone")
	old_password := ctx.PostForm("old_password")
	new_password := ctx.PostForm("new_password")

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
	if len(new_password) < 6 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code":    422,
			"message": "密码不能少于6位",
		})
		return
	}

	//判断用户名和手机号是否存在
	var user models.User
	u.db.Model(&models.User{}).Where("name = ? AND telephone = ?", name, telephone).First(&user)
	if user.ID == 0 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code":    422,
			"message": "用户名或手机号错误",
		})
		return
	}

	//判断旧密码是否正确
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(old_password)); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code":    422,
			"message": "旧密码错误",
		})
		return
	}

	//修改密码
	hashednewPassword, err := bcrypt.GenerateFromPassword([]byte(new_password), bcrypt.DefaultCost)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code":    500,
			"message": "密码加密错误",
		})
		return
	}

	result := u.db.Model(&models.User{}).Where("name = ? AND telephone = ?", name, telephone).Update("password", string(hashednewPassword))
	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "密码更新失败",
		})
		return
	}
	if result.RowsAffected == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "未找到匹配的用户",
		})
		return
	}

	//返回结果
	ctx.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "修改密码成功了",
	})
}
