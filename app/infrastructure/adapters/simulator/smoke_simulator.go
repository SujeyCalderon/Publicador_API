package simulator

import (
	"math/rand"
	"time"
)

// SimulateSmoke simula una lectura del sensor de humo.
// Devuelve el nivel de humo (por ejemplo, entre 0 y 100) y un indicador de alarma si supera un umbral.
func SimulateSmoke() (float64, bool) {
	rand.Seed(time.Now().UnixNano())
	smokeLevel := rand.Float64() * 100 // Ejemplo: nivel entre 0 y 100
	// Define un umbral para activar la alarma.
	alarm := smokeLevel > 50
	return smokeLevel, alarm
}
