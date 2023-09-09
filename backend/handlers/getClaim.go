package handlers

import (
	"backend/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 获取用户信息和claim
func (u *UserController) GetUserClaimByTime(c *gin.Context) {
	// 上面我们在JWTAuth()中间中将'claims'写入到gin.Context的指针对象中，因此在这里可以将之解析出来
	claims := c.MustGet("claims").(*utils.CustomClaims)
	if claims != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  0,
			"message": "token有效",
			"data":    claims,
		})
	}
}
