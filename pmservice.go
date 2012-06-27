package libclient3

// this struct pointer implements Service.
type PMService struct {
	connection *Connection

}

// connects to the ProcessManager service.
func ConnectPM() (pm *PMService, err error) {

	// setup ProcessManager event handlers.
	if PMEvents == nil {
		CreatePMEvents()
	}

	// connect to ProcessManager.
	conn, err := Connect("/system/socket/ProcessSocket")
	if err == nil {
		pm = &PMService{conn}
		conn.service = pm
	}
	return
}

func (*PMService) handleEvent(command string, params map[string]interface{}) {
	if PMEvents[command] != nil {
		PMEvents[command](params)
	}
}
