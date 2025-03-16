package simulator

import (
	"math/rand"
	"time"
)


func SimulateDoor() bool {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(2) == 1
}
