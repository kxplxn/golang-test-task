package rabbitmq

import (
	"fmt"

	amqp "github.com/rabbitmq/amqp091-go"
)

const connStr = "amqp://user:password@localhost:7001/"

type rabbitmq struct {
	Channel *amqp.Channel
	Queue   amqp.Queue
}

var instance *rabbitmq

func GetChannelAndQueue() (*rabbitmq, error) {
	if instance == nil {
		conn, err := amqp.Dial(connStr)
		if err != nil {
			return nil, fmt.Errorf("error dialing rabbitmq: %s", err)
		}

		ch, err := conn.Channel()
		if err != nil {
			return nil, fmt.Errorf("error opening rabbitmq channel: %s", err)
		}

		q, err := ch.QueueDeclare("message", false, false, false, false, nil)
		if err != nil {
			return nil, fmt.Errorf("error declaring rabbitmq queue: %s", err)
		}

		instance = &rabbitmq{Channel: ch, Queue: q}

		return instance, nil
	} else {
		return instance, nil
	}
}
