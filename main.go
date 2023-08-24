package main

import (
	"NearbySocial/gormfunc"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

var clients = make(map[*websocket.Conn]bool)

func asciiArrayToInt(asciiArray []uint8) int {
	result := 0
	for _, digit := range asciiArray {
		// 将ASCII码转换为数字
		if digit >= 48 && digit <= 57 {
			result = result*10 + int(digit-48)
		} else {
			fmt.Println("Invalid ASCII code:", digit)
		}
	}
	return result
}

func main() {
	db, err := gormfunc.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		sqlDB, _ := db.DB()
		sqlDB.Close()
	}()

	//gin ws暂时放这，后面再改
	r := gin.Default()

	r.GET("/ws", func(c *gin.Context) {
		upgrader := websocket.Upgrader{}
		upgrader.CheckOrigin = func(r *http.Request) bool {
			//允许所有连接
			return true
		}
		ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			log.Println(err)
			return
		}

		clients[ws] = true
		fmt.Println("new client connected")

		for {
			_, message, err := ws.ReadMessage()
			if err != nil {
				log.Println(err)
				break
			}

			//仅实现了根据id查询用户，id为int
			//intValue := int(message[0])<<24 | int(message[1])<<16 | int(message[2])<<8 | int(message[3])
			intValue := asciiArrayToInt(message)
			person, err := gormfunc.GetUserByID(db, intValue)
			if err != nil {
				log.Println(err)
				break
			}
			message, err = json.Marshal(person)
			err = ws.WriteMessage(websocket.TextMessage, []byte(message))
			if err != nil {
				log.Println("write error:", err)
				ws.Close()
				delete(clients, ws)
			}
		}
	})
	r.Run()
}
