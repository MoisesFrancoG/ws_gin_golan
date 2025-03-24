package domain

import (
	"log"
	"time"

	"github.com/gorilla/websocket"
)

type Session struct {
	conn      *websocket.Conn
	SessionID string
}

func NewSession(conn *websocket.Conn, userID string) *Session {
	return &Session{
		conn:      conn,
		SessionID: userID,
	}
}

func (s *Session) StartHandling() {
	log.Println(s.SessionID)
	s.readPump()
	//s.writePump()
}

func (s *Session) readPump() {
	defer s.conn.Close()

	go s.writePump()

	for {

		messageType, p, err := s.conn.ReadMessage()

		if err != nil {
			if websocket.IsUnexpectedCloseError(
				err,
				websocket.CloseGoingAway,
				websocket.CloseAbnormalClosure,
			) {
				log.Printf("Error %v", err)
				break
			}
		}

		if messageType != -1 {
			log.Printf("Recibido: %s, de tipo %d", p, messageType)

			s.broadcast(1, p)
		}

		time.Sleep(17 * time.Millisecond)
	}

	//select {}
}

func (s *Session) writePump() {
	defer func() {
		s.conn.Close()
	}()

	for {

		messageType := websocket.TextMessage
		message := []byte("Message from server")

		err := s.conn.WriteMessage(messageType, message)

		if err != nil {
			log.Println("Write error: ", err)
			break
		}

		time.Sleep(10 * time.Second)
	}

	select {}
}

func (s *Session) broadcast(messageType int, payloadbyte []byte) {
	err := s.conn.WriteMessage(messageType, payloadbyte)
	if err != nil {
		log.Println("Broadcast error: ", err)
	}
}
