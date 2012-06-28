package libclient3

var PMEvents map[string]func(*PMService, map[string]interface{})

func CreatePMEvents() {
	PMEvents = map[string]func(*PMService, map[string]interface{}){
		"ping": pingHandler,
	}
}

// keep the connection alive.
func pingHandler(pm *PMService, _ map[string]interface{}) {
	pm.connection.Send("pong", nil)
}
