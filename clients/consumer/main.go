// consumer.go
package main

import (
	"fmt"
	"log"
	"net/url"
	"sync"

	//"time"

	"github.com/gorilla/websocket"
)

var (
	messages []string
	mutex    sync.Mutex
)

func consumeMessages() {
	serverURL := url.URL{Scheme: "ws", Host: "localhost:8080", Path: "/ws/handshake", RawQuery: "user_id=consumer"}
	conn, _, err := websocket.DefaultDialer.Dial(serverURL.String(), nil)
	if err != nil {
		log.Fatal("Error connecting to WebSocket server:", err)
	}
	defer conn.Close()

	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("Read error:", err)
			return
		}
		mutex.Lock()
		messages = append(messages, string(message))
		mutex.Unlock()
		fmt.Println("Received:", string(message))
	}
}

func displayMessages() {
	for {
		//
		mutex.Lock()
		fmt.Println("\nMessages Received:")
		for _, msg := range messages {
			fmt.Println(msg)
		}
		mutex.Unlock()
	}
}

func main() {
	go consumeMessages()
	//go displayMessages()
	select {}
}
