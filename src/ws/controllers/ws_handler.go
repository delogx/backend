package controllers

import (
	"backend/src/ws/types"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func WSHandler(c *gin.Context, rm *types.RoomManager) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println("Failed to upgrade connection to Websocket: ", err)
		return
	}

	log.Println("Client connected")

	defer func() {
		log.Println("Client disconnected")
		conn.Close()
	}()

	roomID := c.Param("roomId")
	rm.JoinRoom(roomID, conn)

	go func() {
		for {
			messageType, message, err := conn.ReadMessage()
			if err != nil {
				log.Println("Error reading message: ", err)
				break
			}

			switch messageType {
			case websocket.TextMessage:
				// Handle text messages
				log.Println("Received text message:", string(message))
				if string(message) == "ping" {
					if err := conn.WriteMessage(websocket.TextMessage, []byte("pong")); err != nil {
						log.Println("Error sending pong message: ", err)
						return
					}
				} else {
					log.Println("Received text message: ", string(message))
				}
			case websocket.CloseMessage:
				// Handle close messages
				log.Println("Received close message")
				rm.LeaveRoom(roomID, conn)
				return
			}
		}
	}()

	for {
	}
}
