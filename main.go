package main

import (
	"example.com/discord_bot_backend_boiler/controller"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/ping", controller.ResponsePong)

	router.Run("localhost:1010")
}
