package domain
import "time"

type MotionSensor struct {
	ID             string  `json:"id"`              
	Timestamp time.Time `json:"timestamp"`    
	MotionDetected bool    `json:"motion_detected"` 
	Intensity      float64 `json:"intensity"`       
	DeviceToken string    `json:"deviceToken"` 
}
