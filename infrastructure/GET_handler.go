package infrastructure

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

func WSHandler(ctx *gin.Context) {

	var upgrader websocket.Upgrader

	if websocket.IsWebSocketUpgrade(ctx.Request) {
		upgrader = websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		}
	} else {
		ctx.JSON(http.StatusBadRequest, gin.H{"Error": "Cliente no aceptado"})
	}
	/* var upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	} */

	conn, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)

	if err != nil {
		log.Printf("Error de conexión: %v", err)
		return
	}

	defer conn.Close()

	log.Println("Cliente conectado")

	for {

		messageType, p, err := conn.ReadMessage()

		if err != nil {
			log.Printf("Error en la lectura: %v", err)
			break
		}

		log.Printf("Recibido: %s, de tipo %d", p, messageType)

		conn.WriteMessage(1, p)

		time.Sleep(16 * time.Millisecond)
	}

	log.Println("Cliente desconectado")

}
