package domain
import "time"

type LightSensor struct {
	ID         string  `json:"id"`         
	Timestamp time.Time `json:"timestamp"`
	Luminosity float64 `json:"luminosity"` 
	DeviceToken string    `json:"deviceToken"` 
}
