package utils

import (
	"backend/models"
	"math"
)

func DistanceBetweenUsers(user1 *models.User, user2 *models.User) float64 {

	//由于geohash.Convert函数返回的是弧度值，这里先转换为角度
	lat1, lon1 := user1.Latitude*math.Pi/180, user1.Longitude*math.Pi/180
	lat2, lon2 := user2.Latitude*math.Pi/180, user2.Longitude*math.Pi/180

	R := 6371.0 //地球半径常量

	//使用haversine公式计算球面距离
	dlon := lon2 - lon1
	dlat := lat2 - lat1
	a := math.Pow(math.Sin(dlat/2), 2) + math.Cos(lat1)*math.Cos(lat2)*math.Pow(math.Sin(dlon/2), 2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	distance := R * c

	return distance
}
