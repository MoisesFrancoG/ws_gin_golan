package infrastructure

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

func handleWSClient(ctx *gin.Context) {

	var upgr websocket.Upgrader

	if websocket.IsWebSocketUpgrade(ctx.Request) {
		upgr = websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		}
	} else {
		ctx.JSON(400, gin.H{"Error": "Esta ruta necesita un cliente WebSocket"})
	}

	/* upgr := websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	} */

	conn, err := upgr.Upgrade(ctx.Writer, ctx.Request, nil)

	if err != nil {
		log.Println("Error durante upgrade")
		return
	}

	defer conn.Close()

	log.Println("Cliente conectado")

	for {
		messageType, p, err := conn.ReadMessage()

		if err != nil {
			log.Println("Error al leer el mensaje: ", err)
		}

		log.Printf("Recibido: %s de tipo %d\n", p, messageType)

		err = conn.WriteMessage(messageType, p)

		if err != nil {
			log.Println("Error al enviar mensaje")
			break
		}
	}

	log.Println("Desconectado")

}
