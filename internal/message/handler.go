package message

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"

	"twitch_chat_analysis/pkg/rabbitmq"

	"github.com/gin-gonic/gin"
	amqp "github.com/rabbitmq/amqp091-go"
)

type message struct {
	Sender   string `json:"sender"`
	Receiver string `json:"receiver"`
	Message  string `json:"message"`
}

func Handle(c *gin.Context) {
	var msg message
	if err := json.NewDecoder(c.Request.Body).Decode(&msg); err != nil {
		log.Printf("error reading request body: %s", err)
		c.String(http.StatusInternalServerError, "")
		return
	}
	if msg.Sender == "" || msg.Receiver == "" || msg.Message == "" {
		c.String(http.StatusBadRequest, "")
		return
	}

	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		log.Printf("error reading request body: %s", err)
		c.String(http.StatusInternalServerError, "")
		return
	}

	rmq, err := rabbitmq.Get()
	if err != nil {
		log.Printf("error getting rabbitmq channel: %s", err)
		c.String(http.StatusInternalServerError, "")
		return
	}

	err = rmq.Channel.PublishWithContext(
		context.Background(),
		"",
		rmq.Queue.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)
	if err != nil {
		log.Printf("error publishing to rabbitmq: %s", err)
		c.String(http.StatusInternalServerError, "")
		return
	}
}
