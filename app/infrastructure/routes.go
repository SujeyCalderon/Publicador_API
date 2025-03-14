package infrastructure

import (
	"net/http"
)

func SetupRoutes(deps *Dependencies) {
	http.HandleFunc("/sensor/door", deps.DoorHandler.HandleDoor)
	http.HandleFunc("/sensor/light", deps.LightHandler.HandleLight)
	http.HandleFunc("/sensor/motion", deps.MotionHandler.HandleMotion)
	http.HandleFunc("/sensor/smoke", deps.SmokeHandler.HandleSmoke)
}
