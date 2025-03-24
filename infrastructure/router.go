package infrastructure

import "github.com/gin-gonic/gin"

func Routes(engine *gin.Engine) {
	ws_group := engine.Group("ws")

	ws_group.GET("handshake", WSHandler)

}
