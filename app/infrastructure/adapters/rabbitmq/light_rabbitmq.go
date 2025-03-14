package rabbitmq

import (
	"encoding/json"
	"log"

	"github.com/SUJEY/PUBLICADOR/app/domain"
	"github.com/streadway/amqp"
)

type LightRabbitMQ struct {
	conn    *amqp.Connection
	channel *amqp.Channel
	queue   string
}

func NewLightRabbitMQ(conn *amqp.Connection, queue string) (*LightRabbitMQ, error) {
	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}
	_, err = ch.QueueDeclare(
		queue,
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return nil, err
	}
	return &LightRabbitMQ{
		conn:    conn,
		channel: ch,
		queue:   queue,
	}, nil
}

func (r *LightRabbitMQ) Publish(sensor domain.LightSensor) error {
	body, err := json.Marshal(sensor)
	if err != nil {
		return err
	}
	err = r.channel.Publish(
		"",
		r.queue,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)
	if err != nil {
		return err
	}
	log.Printf("Published light sensor data to RabbitMQ: %+v\n", sensor)
	return nil
}
