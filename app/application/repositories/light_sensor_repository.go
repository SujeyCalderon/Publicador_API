package repositories

import (
	"github.com/SUJEY/PUBLICADOR/app/domain"
	"github.com/SUJEY/PUBLICADOR/app/infrastructure/adapters/rabbitmq"
	"github.com/streadway/amqp"
)

type LightSensorRepository interface {
	Save(sensor domain.LightSensor) error
}

type LightSensorRepositoryImpl struct {
	publisher *rabbitmq.LightRabbitMQ
}

func NewLightSensorRepository(conn *amqp.Connection) (LightSensorRepository, error) {
	publisher, err := rabbitmq.NewLightRabbitMQ(conn, "light")
	if err != nil {
		return nil, err
	}
	return &LightSensorRepositoryImpl{publisher: publisher}, nil
}

func (repo *LightSensorRepositoryImpl) Save(sensor domain.LightSensor) error {
	return repo.publisher.Publish(sensor)
}
