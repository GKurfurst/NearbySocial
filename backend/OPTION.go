package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// CORSMiddleware 中间件处理CORS问题
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		origin := c.Request.Header.Get("Origin")
		if c.Request.Method == "OPTIONS" {
			// 设置header
			c.Header("Access-Control-Allow-Origin", origin) // 允许的域名，如 http://localhost:5173
			c.Header("Access-Control-Allow-Methods", "GET,POST,PUT,DELETE,PATCH,OPTIONS")
			c.Header("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
			c.Header("Access-Control-Allow-Credentials", "true")
			c.Header("Access-Control-Max-Age", "86400") // 设置预检请求的结果的缓存时间，24小时

			c.AbortWithStatus(http.StatusNoContent)
			return
		}
		c.Next()
	}
}

func main() {
	app := gin.Default()
	app.Use(CORSMiddleware()) // 使用CORS中间件

	// 定义路由
	app.GET("/route", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "This is cross-origin route1",
		})
	})

	// 运行服务
	app.Run("localhost:8080")
}
