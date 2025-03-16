package infrastructure

import (
	"log"

	"github.com/SUJEY/PUBLICADOR/app/application/repositories"
	"github.com/SUJEY/PUBLICADOR/app/application/services"
	"github.com/SUJEY/PUBLICADOR/app/application/usecase"
	"github.com/SUJEY/PUBLICADOR/app/infrastructure/adapters/rest_handler"
	"github.com/streadway/amqp"
)

// Dependencies agrupa todos los handlers de la API.
type Dependencies struct {
	DoorHandler   *rest_handler.DoorHandler
	LightHandler  *rest_handler.LightHandler
	MotionHandler *rest_handler.MotionHandler
	SmokeHandler  *rest_handler.SmokeHandler
}

func InitDependencies() *Dependencies {
	// Conexión a RabbitMQ (ajusta la URL de conexión según tu entorno)
	conn, err := amqp.Dial("amqp://Sujey:calmar58@34.194.130.32/")
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}
	// Nota: considera guardar la conexión para cerrarla en una función de cleanup si lo requieres.

	// Inicializar repositorios pasando la conexión
	doorRepo, err := repositories.NewDoorSensorRepository(conn)
	if err != nil {
		log.Fatalf("Failed to initialize door repository: %v", err)
	}
	lightRepo, err := repositories.NewLightSensorRepository(conn)
	if err != nil {
		log.Fatalf("Failed to initialize light repository: %v", err)
	}
	motionRepo, err := repositories.NewMotionSensorRepository(conn)
	if err != nil {
		log.Fatalf("Failed to initialize motion repository: %v", err)
	}
	smokeRepo, err := repositories.NewSmokeSensorRepository(conn)
	if err != nil {
		log.Fatalf("Failed to initialize smoke repository: %v", err)
	}

	// Crear servicios
	doorService := services.NewDoorService(doorRepo)
	lightService := services.NewLightService(lightRepo)
	motionService := services.NewMotionService(motionRepo)
	smokeService := services.NewSmokeService(smokeRepo)

	// Casos de uso
	doorUC := usecase.NewDoorUseCase(doorService)
	lightUC := usecase.NewLightUseCase(lightService)
	motionUC := usecase.NewMotionUseCase(motionService)
	smokeUC := usecase.NewSmokeUseCase(smokeService)

	// Handlers REST
	return &Dependencies{
		DoorHandler:   rest_handler.NewDoorHandler(doorUC),
		LightHandler:  rest_handler.NewLightHandler(lightUC),
		MotionHandler: rest_handler.NewMotionHandler(motionUC),
		SmokeHandler:  rest_handler.NewSmokeHandler(smokeUC),
	}
}
