package domain
import "time"

type DoorSensor struct {
	ID        string `json:"id"`        
	Timestamp time.Time `json:"timestamp"` 
	IsOpen    bool   `json:"is_open"`   
	DeviceToken string    `json:"deviceToken"` 
}
