package routes

import (
	"gin-service/controller"

	"github.com/gin-gonic/gin"
)

func PingRoute(router *gin.Engine) {
	router.GET("/ping", controller.GetPong)
}
