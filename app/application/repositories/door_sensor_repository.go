package repositories

import (
	"github.com/SUJEY/PUBLICADOR/app/domain"
	"github.com/SUJEY/PUBLICADOR/app/infrastructure/adapters/rabbitmq"
	"github.com/streadway/amqp"
)

type DoorSensorRepository interface {
	Save(sensor domain.DoorSensor) error
}

type DoorSensorRepositoryImpl struct {
	publisher *rabbitmq.DoorRabbitMQ
}

func NewDoorSensorRepository(conn *amqp.Connection) (DoorSensorRepository, error) {
	publisher, err := rabbitmq.NewDoorRabbitMQ(conn, "door")
	if err != nil {
		return nil, err
	}
	return &DoorSensorRepositoryImpl{publisher: publisher}, nil
}

func (repo *DoorSensorRepositoryImpl) Save(sensor domain.DoorSensor) error {
	return repo.publisher.Publish(sensor)
}
