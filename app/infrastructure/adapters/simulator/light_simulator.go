package simulator

import (
	"math/rand"
	"time"
)

func SimulateLight() float64 {
	rand.Seed(time.Now().UnixNano())
	return float64(rand.Intn(1001))
}
