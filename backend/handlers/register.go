package handlers

import (
	"backend/models"
	"backend/utils"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"strconv"
)

func (u *UserController) UserRegister(ctx *gin.Context) {
	//获取参数
	var requestBody struct {
		Username  string `json:"username"`
		Telephone string `json:"telephone"`
		Password  string `json:"password"`
	}
	if err := ctx.ShouldBindJSON(&requestBody); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code":    422,
			"message": "请求数据无效",
		})
		return
	}
	name := requestBody.Username
	telephone := requestBody.Telephone
	password := requestBody.Password

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

	//判断手机号是否存在
	u.db.Model(&models.User{}).Where("telephone = ?", telephone).First(&user)
	if user.ID != 0 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code":    422,
			"message": "手机号已存在",
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

	var id int64
	for {
		id = utils.GenerateUserId()
		u.db.Model(&models.User{}).Where("user_id = ?", id).First(&user)
		if user.ID == 0 {
			break
		}
	}

	newUser := models.User{
		Name:      name,
		UserId:    strconv.FormatInt(id, 10),
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
