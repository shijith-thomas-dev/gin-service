package routes

import (
	"gin-service/controller"

	"github.com/gin-gonic/gin"
)

func UserRoute(router *gin.Engine) {
	router.GET("/", controller.GetUsers)
	router.POST("/", controller.CreateUser)
	router.PUT("/:id", controller.UpdateUser)
	router.DELETE("/:id", controller.DeleteUser)
}
