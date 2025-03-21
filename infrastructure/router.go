package infrastructure

import (
	"github.com/gin-gonic/gin"
)

func SetRoutes(engine *gin.Engine) {
	group := engine.Group("/ws")

	group.GET("/handshake", handleWSClient)
}
