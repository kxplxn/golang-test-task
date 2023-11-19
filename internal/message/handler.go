package message

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	amqp "github.com/rabbitmq/amqp091-go"
)

const connStr = "amqp://user:password@localhost:7001/"

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

	conn, err := amqp.Dial(connStr)
	if err != nil {
		log.Printf("error dialing rabbitmq: %s", err)
		c.String(http.StatusInternalServerError, "")
		return
	}

	ch, err := conn.Channel()
	if err != nil {
		log.Printf("error opening rabbitmq channel: %s", err)
		c.String(http.StatusInternalServerError, "")
		return
	}

	q, err := ch.QueueDeclare("message", false, false, false, false, nil)
	if err != nil {
		log.Printf("error declaring rabbitmq queue: %s", err)
		c.String(http.StatusInternalServerError, "")
		return
	}

	err = ch.PublishWithContext(
		context.Background(),
		"",
		q.Name,
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
