package main

import (
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/gorilla/websocket"
	"log"
	"time"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func main() {
	r := gin.Default()
	r.GET("/chat/:friendId", handleWebSocket)
	r.Run(":8080")
}

func handleWebSocket(c *gin.Context) {
	friendId := c.Param("friendId")
	userId := c.Query("userId")

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
	pubsub := redisClient.Subscribe(c, friendId+":"+userId)
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
			log.Println("Failed to receive message:", err)
			break
		}

		err = conn.WriteMessage(websocket.TextMessage, []byte(msg.Payload))
		if err != nil {
			log.Println("Failed to send message:", err)
			break
		}
	}
}
