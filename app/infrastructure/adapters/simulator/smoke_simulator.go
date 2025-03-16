package simulator

import (
	"math/rand"
	"time"
)


func SimulateSmoke() (float64, bool) {
	rand.Seed(time.Now().UnixNano())
	smokeLevel := rand.Float64() * 100 

	alarm := smokeLevel > 50
	return smokeLevel, alarm
}
