package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Pong struct {
	Ping string `json:"ping"`
}

func GetPong(c *gin.Context) {
	ping := Pong{
		Ping: "Pong",
	}
	c.JSON(http.StatusOK, &ping)
}
