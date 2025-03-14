package simulator

import (
	"math/rand"
	"time"
)

// SimulateMotion simula la detección de movimiento y la intensidad.
// Devuelve un valor booleano y un valor de intensidad (entre 0 y 1).
func SimulateMotion() (bool, float64) {
	rand.Seed(time.Now().UnixNano())
	motionDetected := rand.Intn(2) == 1
	intensity := 0.0
	if motionDetected {
		intensity = rand.Float64()
	}
	return motionDetected, intensity
}
