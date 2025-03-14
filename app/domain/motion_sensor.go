package domain
import "time"

type MotionSensor struct {
	ID             string  `json:"id"`              
	Timestamp time.Time `json:"timestamp"`    
	MotionDetected bool    `json:"motion_detected"` // true si se detectó movimiento.
	Intensity      float64 `json:"intensity"`       // Intensidad del movimiento (por ejemplo, de 0 a 1).
}
