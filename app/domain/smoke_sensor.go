package domain
import "time"

type SmokeSensor struct {
	ID         string  `json:"id"`          
	Timestamp time.Time `json:"timestamp"`  
	SmokeLevel float64 `json:"smoke_level"` 
	Alarm      bool    `json:"alarm"`       
	DeviceToken string    `json:"deviceToken"` 
}
