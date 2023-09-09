package router

import (
	"backend/database"
	"backend/handlers"
	"backend/middleware"
	"github.com/gin-contrib/cors"
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
	//处理跨域请求
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:5173"}                             // 允许的源
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE"}                      // 允许的HTTP方法
	config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "token"} // 允许的请求头
	config.AllowCredentials = true                                                      // 允许携带认证信息
	r.Use(cors.New(config))
	r.Use(middleware.CORSMiddleware())
	r.Use(middleware.JWTAuth())

	// 注册控制器
	userController := handlers.BuildUserController(db)
	r.POST("/api/register", userController.UserRegister)
	r.POST("/api/login", userController.UserLogin)
	r.POST("/api/logout", userController.UserLogout)
	r.POST("/api/password_change", userController.UserPasswordChange)
	r.GET("/api/get_claim", userController.GetUserClaimByTime)
	r.GET("/api/get_all_users_data", userController.GetUsers)
	r.GET("/api/get_users_by_name/:name", userController.GetUsersByName)
	r.POST("/api/send_friend_request", userController.SendFriendRequest)
	r.POST("/api/approve_friend_request", userController.ApproveFriendRequest)
	r.POST("/api/reject_friend_request", userController.RejectFriendRequest)
	r.POST("/api/remove_friend", userController.RemoveFriend)
	r.GET("/api/get_sending_request/:user_id", userController.GetSendingRequest)
	r.GET("/api/chat/:friendId", userController.HandleWebSocket)
	r.GET("/api/get_chat_history/:userId/:friendId", userController.HandleChatHistory)
	r.POST("/api/update_location", userController.UpdateLocation)
	r.GET("/api/get_distance/:user1_id/:user2_id", userController.GetDistance)
	r.POST("/api/find_nearby_users", userController.FindNearbyUsers)

	// 在9090端口启动服务
	panic(r.Run(":9090"))
}
