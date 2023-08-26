package router

import (
	"backend/database"
	"backend/handlers"
	"backend/middleware"
	"github.com/gin-gonic/gin"
)

func InitRouter() {

	// 获取初始化的数据库
	db := database.InitDB()
	// 延迟关闭数据库
	defer db.Close()

	// 创建一个默认的路由引擎
	r := gin.Default()
	r.Use(middleware.JWTAuth())

	// 注册控制器
	userController := handlers.BuildUserController(db)
	r.POST("/register", userController.UserRegister)
	r.POST("/login", userController.UserLogin)
	r.POST("/password_change", userController.UserPasswordChange)
	r.GET("/getclaim", userController.GetUserClaimByTime)

	// 在9090端口启动服务
	panic(r.Run(":9090"))
}
