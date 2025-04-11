package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/url"
	"time"

	"github.com/gorilla/websocket"
)

func startClient(userID string) {
	serverURL := url.URL{Scheme: "ws", Host: "localhost:8080", Path: "/ws/handshake", RawQuery: fmt.Sprintf("user_id=%s", userID)}
	conn, _, err := websocket.DefaultDialer.Dial(serverURL.String(), nil)
	if err != nil {
		log.Fatal("Error connecting to WebSocket server:", err)
	}
	defer conn.Close()

	for {
		message := fmt.Sprintf("Message from %s at %v", userID, time.Now().Format(time.RFC3339))
		if err := conn.WriteMessage(websocket.TextMessage, []byte(message)); err != nil {
			log.Println("Write error:", err)
			return
		}
		fmt.Println("Sent:", message)
		time.Sleep(time.Duration(rand.Intn(3)+1) * time.Second) 
	}
}

func main() {
	go startClient("client2")
	select {} 
}
