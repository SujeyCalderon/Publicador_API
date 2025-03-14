package domain
import "time"

type SmokeSensor struct {
	ID         string  `json:"id"`          
	Timestamp time.Time `json:"timestamp"`  
	SmokeLevel float64 `json:"smoke_level"` // Nivel de concentración de humo 
	Alarm      bool    `json:"alarm"`       // true si el nivel de humo activa la alarma.
}
