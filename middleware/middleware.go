package middleware

import (
	"backend/utils"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strings"
)

// 定义一个JWTAuth的中间件
func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		//过滤是否验证token， login结构直接放行，这里为了简单起见，直接判断路径中是否带login和register，携带就直接放行
		if strings.Contains(c.Request.RequestURI, "login") || strings.Contains(c.Request.RequestURI, "register") {
			return
		}

		token := c.Request.Header.Get("token")
		if token == "" {
			c.JSON(http.StatusOK, gin.H{
				"status": -1,
				"msg":    "请求未携带token，无权限访问",
				"data":   nil,
			})
			c.Abort()
			return
		}

		log.Print("get token: ", token)

		// 初始化一个JWT对象实例，并根据结构体方法来解析token
		j := utils.NewJWT()
		// 解析token中包含的相关信息(有效载荷)
		claims, err := j.ParserToken(token)

		if err != nil {
			// 	与token无关的其他错误
			c.JSON(http.StatusOK, gin.H{
				"status": -1,
				"msg":    err.Error(),
				"data":   nil,
			})
			c.Abort()
			return
		}

		// 将解析后的有效载荷claims重新写入gin.Context引用对象中
		c.Set("claims", claims)

	}
}
