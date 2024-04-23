package types

import (
	"log"
	"sync"

	"github.com/gorilla/websocket"
)

type RoomManager struct {
	rooms map[string]map[*websocket.Conn]struct{}
	mu    sync.Mutex
}

func NewRoomManager() *RoomManager {
	return &RoomManager{
		rooms: make(map[string]map[*websocket.Conn]struct{}),
	}
}

func (rm *RoomManager) JoinRoom(roomID string, conn *websocket.Conn) {
	rm.mu.Lock()
	defer rm.mu.Unlock()

	if _, ok := rm.rooms[roomID]; !ok {
		rm.rooms[roomID] = make(map[*websocket.Conn]struct{})
	}
	rm.rooms[roomID][conn] = struct{}{}
}

func (rm *RoomManager) LeaveRoom(roomID string, conn *websocket.Conn) {
	rm.mu.Lock()
	defer rm.mu.Unlock()

	if _, ok := rm.rooms[roomID]; !ok {
		return
	}
	delete(rm.rooms[roomID], conn)
	if len(rm.rooms[roomID]) == 0 {
		delete(rm.rooms, roomID)
	}
}

func (rm *RoomManager) Broadcast(roomID string, messageType int, message []byte, currentConn *websocket.Conn) {
	rm.mu.Lock()
	defer rm.mu.Unlock()
	clients, ok := rm.rooms[roomID]
	if !ok {
		return
	}
	for conn := range clients {
		if conn == currentConn {
			continue
		}
		if err := conn.WriteMessage(messageType, message); err != nil {
			log.Println("Error broadcasting message to room", roomID, ":", err)
		}
	}
}
