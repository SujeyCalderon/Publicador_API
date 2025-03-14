package repositories

import (
	"github.com/SUJEY/PUBLICADOR/app/domain"
	"github.com/SUJEY/PUBLICADOR/app/infrastructure/adapters/rabbitmq"
	"github.com/streadway/amqp"
)

type MotionSensorRepository interface {
	Save(sensor domain.MotionSensor) error
}

type MotionSensorRepositoryImpl struct {
	publisher *rabbitmq.MotionRabbitMQ
}

func NewMotionSensorRepository(conn *amqp.Connection) (MotionSensorRepository, error) {
	publisher, err := rabbitmq.NewMotionRabbitMQ(conn, "motion")
	if err != nil {
		return nil, err
	}
	return &MotionSensorRepositoryImpl{publisher: publisher}, nil
}

func (repo *MotionSensorRepositoryImpl) Save(sensor domain.MotionSensor) error {
	return repo.publisher.Publish(sensor)
}
