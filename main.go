package main

import (
	"log"
	"net/http"

	"github.com/SUJEY/PUBLICADOR/app/infrastructure"
)

func main() {
	// Inicializar dependencias
	deps := infrastructure.InitDependencies()

	// Configurar rutas
	infrastructure.SetupRoutes(deps)

	log.Println("API de productor corriendo en el puerto 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Error iniciando el servidor: %v", err)
	}
}
