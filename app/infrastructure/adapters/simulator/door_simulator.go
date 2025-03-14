package simulator

import (
	"math/rand"
	"time"
)

// SimulateDoor simula el estado de la puerta (abierta o cerrada).
func SimulateDoor() bool {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(2) == 1
}
