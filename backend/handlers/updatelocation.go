package handlers

import (
	"backend/models"
	"github.com/gin-gonic/gin"
	"github.com/mmcloughlin/geohash"
	"net/http"
)

func (u *UserController) UpdateLocation(ctx *gin.Context) {

	//获取参数
	var requestBody struct {
		UserId    string  `json:"user_id"`
		Longitude float64 `json:"longitude"`
		Latitude  float64 `json:"latitude"`
	}
	if err := ctx.ShouldBindJSON(&requestBody); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code":    422,
			"message": "请求数据无效",
		})
		return
	}

	userID := requestBody.UserId
	longitude := requestBody.Longitude
	latitude := requestBody.Latitude

	//验证数据
	var user models.User
	if err := u.db.Where("user_id = ?", userID).First(&user).Error; err != nil {
		ctx.JSON(400, gin.H{"error": "User not found"})
		return
	}

	//更新数据库位置信息
	user.Longitude = longitude
	user.Latitude = latitude
	user.GeoHash = geohash.Encode(latitude, longitude)
	u.db.Save(&user)

	//返回结果
	ctx.JSON(http.StatusOK, gin.H{
		"message": "用户位置信息已更新",
	})

}
