package controller

import (
	"net/http"

	"example.com/discord_bot_backend_boiler/model"
	"github.com/gin-gonic/gin"
)

func ResponsePong(c *gin.Context) {

	pong := model.Pong{
		Text: "Pong",
	}

	c.IndentedJSON(http.StatusOK, pong)
}
