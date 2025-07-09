package handlers

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var (
	clients   = make(map[*websocket.Conn]bool)
	broadcast = make(chan []byte)

	upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)

func WebSocketHandler(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Error upgrading to WebSocket:", err)
		return
	}
	defer ws.Close()

	clients[ws] = true
	log.Println("Client connected:", ws.RemoteAddr())

	for {
		var dummy string
		if err := ws.ReadJSON(&dummy); err != nil {
			delete(clients, ws)
			log.Println("Client disconnected:", ws.RemoteAddr())
			break
		}
	}

}

func StartBroadcast() {
	for {
		msg := <-broadcast
		for client := range clients {
			err := client.WriteJSON(msg)
			if err != nil {
				log.Println("Error sending message to client:", err)
				client.Close()
				delete(clients, client)
			}
		}
	}
}

// func PushNewOrders(Data string) {
// 	broadcast <- Data
// }
