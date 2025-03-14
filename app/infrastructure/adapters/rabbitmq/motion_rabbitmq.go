package rabbitmq

import (
	"encoding/json"
	"log"

	"github.com/SUJEY/PUBLICADOR/app/domain"
	"github.com/streadway/amqp"
)

type MotionRabbitMQ struct {
	conn    *amqp.Connection
	channel *amqp.Channel
	queue   string
}

func NewMotionRabbitMQ(conn *amqp.Connection, queue string) (*MotionRabbitMQ, error) {
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
	return &MotionRabbitMQ{
		conn:    conn,
		channel: ch,
		queue:   queue,
	}, nil
}

func (r *MotionRabbitMQ) Publish(sensor domain.MotionSensor) error {
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
	log.Printf("Published motion sensor data to RabbitMQ: %+v\n", sensor)
	return nil
}
