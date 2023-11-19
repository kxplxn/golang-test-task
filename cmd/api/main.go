package main

import (
	"github.com/gin-gonic/gin"
	"twitch_chat_analysis/internal/message"
)

func main() {
	r := gin.Default()

	r.GET("/test", func(c *gin.Context) {
		c.JSON(200, "worked")
	})

	r.POST("/message", message.Handle)

	r.Run()
}
