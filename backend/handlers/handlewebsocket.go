package handlers

import (
	"backend/models"
	"backend/utils"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"time"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

func (u *UserController) HandleWebSocket(c *gin.Context) {

	//获取参数
	friendId := c.Param("friendId")
	userId := c.Query("userId")

	//判断发送者和接收者是否存在
	var user models.User
	if err := u.db.Preload("Friends").Where("user_id = ?", userId).First(&user).Error; err != nil {
		c.JSON(400, gin.H{"error": "User not found"})
		return
	}

	var friend models.User
	if err := u.db.Preload("Friends").Where("user_id = ?", friendId).First(&friend).Error; err != nil {
		c.JSON(400, gin.H{"error": "Friend not found"})
		return
	}

	if !utils.IsFriend(user, friend) {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Not a friend",
		})
		return
	}

	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println("Failed to upgrade to WebSocket:", err)
		return
	}
	defer conn.Close()

	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	//订阅频道
	pubsub := redisClient.Subscribe(c, userId+":"+friendId)
	defer pubsub.Close()

	// 发送心跳包
	go func() {
		for {
			err := conn.WriteMessage(websocket.PingMessage, nil)
			if err != nil {
				log.Println("Failed to send ping message:", err)
				break
			}
			time.Sleep(30 * time.Second)
		}
	}()

	// 读取客户端发送的消息
	go func() {
		for {
			_, msg, err := conn.ReadMessage()
			if err != nil {
				log.Println("Failed to read message:", err)
				break
			}

			var message models.Message
			err = json.Unmarshal(msg, &message)
			if err != nil {
				log.Println("Failed to unmarshal message:", err)
				return
			}

			if message.SenderID != userId || message.ReceiverID != friendId {
				log.Println("Invalid message:", message)
				continue
			}

			// 将消息的时间戳设置为当前时间
			message.Timestamp = time.Now().Format(time.RFC3339)
			msgWithTimestamp, err := json.Marshal(message)
			if err != nil {
				log.Println("Failed to marshal message:", err)
				continue
			}

			// 将消息存储在Redis的历史列表中
			redisClient.RPush(c, userId+":"+friendId+"_history", string(msgWithTimestamp))

			// 将消息发送给redis频道
			err = redisClient.Publish(c, friendId+":"+userId, string(msg)).Err()
			if err != nil {
				log.Println("Failed to publish message:", err)
				break
			}
		}
	}()

	// 接收redis频道上的消息，发送给客户端
	for {
		msg, err := pubsub.ReceiveMessage(c)
		if err != nil {
			log.Println("Failed to receive messages:", err)
			break
		}

		err = conn.WriteMessage(websocket.TextMessage, []byte(msg.Payload))
		if err != nil {
			log.Println("Failed to send message:", err)
			break
		}
	}
}
