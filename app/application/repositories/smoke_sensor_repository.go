package repositories

import (
	"github.com/SUJEY/PUBLICADOR/app/domain"
	"github.com/SUJEY/PUBLICADOR/app/infrastructure/adapters/rabbitmq"
	"github.com/streadway/amqp"
)

type SmokeSensorRepository interface {
	Save(sensor domain.SmokeSensor) error
}

type SmokeSensorRepositoryImpl struct {
	publisher *rabbitmq.SmokeRabbitMQ
}

func NewSmokeSensorRepository(conn *amqp.Connection) (SmokeSensorRepository, error) {
	publisher, err := rabbitmq.NewSmokeRabbitMQ(conn, "smoke")
	if err != nil {
		return nil, err
	}
	return &SmokeSensorRepositoryImpl{publisher: publisher}, nil
}

func (repo *SmokeSensorRepositoryImpl) Save(sensor domain.SmokeSensor) error {
	return repo.publisher.Publish(sensor)
}
