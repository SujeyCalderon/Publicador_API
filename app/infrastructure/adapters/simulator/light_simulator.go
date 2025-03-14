package simulator

import (
	"math/rand"
	"time"
)

// SimulateLight simula una lectura de luminosidad (por ejemplo, entre 0 y 1000).
func SimulateLight() float64 {
	rand.Seed(time.Now().UnixNano())
	return float64(rand.Intn(1001))
}
