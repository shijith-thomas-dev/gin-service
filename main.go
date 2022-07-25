package main

import (
	"gin-service/config"
	"gin-service/routes"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.New()
	routes.UserRoute(router)
	routes.PingRoute(router)
	os.Setenv("PORT", "8000")
	config.Connect()
	router.Run()
}
