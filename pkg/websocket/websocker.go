package websocket

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type Client struct {
	conn *websocket.Conn
}

var clients = make(map[*Client]bool)
var broadcast = make(chan []byte)

func HandleWebSocket(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Error while upgrading connection:", err)
		return
	}

	client := &Client{conn: conn}
	clients[client] = true

	fmt.Println("New client connected")
	go handleMessages(client)
}

func handleMessages(client *Client) {
	defer func() {
		client.conn.Close()
		delete(clients, client)
		fmt.Println("Client disconnected")
	}()

	for {
		_, message, err := client.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("Error reading message: %v", err)
			}
			break
		}
		BroadcastMessage([]byte(message))
	}
}

func BroadcastMessage(message []byte) {
	broadcast <- message
}

func init() {
	go func() {
		for {
			msg := <-broadcast
			for client := range clients {
				err := client.conn.WriteMessage(websocket.TextMessage, msg)
				if err != nil {
					log.Printf("Error broadcasting to client: %v", err)
					client.conn.Close()
					delete(clients, client)
				}
			}
		}
	}()
}
