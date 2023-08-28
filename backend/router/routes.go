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
	defer func() {
		sqlDB, _ := db.DB()
		sqlDB.Close()
	}()

	// 创建一个默认的路由引擎
	r := gin.Default()
	r.Use(middleware.JWTAuth())

	// 注册控制器
	userController := handlers.BuildUserController(db)
	r.POST("/api/register", userController.UserRegister)
	r.POST("/api/login", userController.UserLogin)
	r.POST("/api/password_change", userController.UserPasswordChange)
	r.GET("/api/get_claim", userController.GetUserClaimByTime)

	// 在9090端口启动服务
	panic(r.Run(":9090"))
}
