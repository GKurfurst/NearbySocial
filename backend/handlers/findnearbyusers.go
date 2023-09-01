package handlers

import (
	"backend/models"
	"backend/utils"
	"github.com/gin-gonic/gin"
	"github.com/mmcloughlin/geohash"
	"net/http"
)

type NearbyUser struct {
	User     models.User `json:"user"`
	Distance float64     `json:"distance"`
}

func (u *UserController) FindNearbyUsers(c *gin.Context) {

	//获取参数
	var requestBody struct {
		UserId string  `json:"user_id"`
		Radius float64 `json:"radius"` //方圆几何，m做单位
	}
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"code":    422,
			"message": "请求数据无效",
		})
		return
	}

	userID := requestBody.UserId
	radius := requestBody.Radius

	//判断用户是否存在
	var user models.User
	if err := u.db.Where("user_id = ?", userID).First(&user).Error; err != nil {
		c.JSON(400, gin.H{"error": "Sender not found"})
		return
	}

	//根据距离选择精确度
	precision := 9 // precision for about 10m
	switch {
	case radius > 10 && radius <= 50:
		precision = 8 // precision for about 50m
	case radius > 50 && radius <= 100:
		precision = 7 // precision for about 150m
	case radius > 100 && radius <= 1000:
		precision = 6 // precision for about 1.2km
	case radius > 1000 && radius <= 10000:
		precision = 5 // precision for about 10km
	case radius > 100000 && radius <= 500000:
		precision = 4 // precision for about 50km
	case radius > 500000 && radius <= 1000000:
		precision = 3 // precision for about 100km
	case radius > 1000000 && radius <= 10000000:
		precision = 2 // precision for about 1000km
	}

	// 找出附近的用户
	neighbors := geohash.Neighbors(user.GeoHash[:precision])
	neighbors_and_self := append(neighbors, user.GeoHash[:precision])

	// 查询数据库
	var nearbyUsers []models.User
	u.db.Where("SUBSTRING(geo_hash, 1, ?) IN (?)", precision, neighbors_and_self).Find(&nearbyUsers)

	// 根据实际距离过滤用户
	var nearbyWithinDistance []NearbyUser
	for _, u := range nearbyUsers {
		distance := utils.DistanceBetweenUsers(&user, &u)
		if distance <= radius && distance > 0 {
			nearbyWithinDistance = append(nearbyWithinDistance, NearbyUser{
				User:     u,
				Distance: distance,
			})
		}
	}

	//返回结果
	c.JSON(http.StatusOK, gin.H{
		"nearby users": nearbyWithinDistance,
	})

}
