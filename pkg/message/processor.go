package message

import (
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
	"twitch_chat_analysis/pkg/rabbitmq"
)

type Processor struct{}

func NewProcessor() Processor { return Processor{} }

func (p Processor) Proccess() (<-chan amqp.Delivery, error) {
	rmq, err := rabbitmq.Get()
	if err != nil {
		return nil, fmt.Errorf("error getting rabbitmq channel: %s", err)
	}
	msgs, err := rmq.Channel.Consume(
		rmq.Queue.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return nil, fmt.Errorf("error consume messages: %s", err)
	}
	return msgs, nil
}
