package simulator

import (
	"math/rand"
	"time"
)


func SimulateMotion() (bool, float64) {
	rand.Seed(time.Now().UnixNano())
	motionDetected := rand.Intn(2) == 1
	intensity := 0.0
	if motionDetected {
		intensity = rand.Float64()
	}
	return motionDetected, intensity
}
