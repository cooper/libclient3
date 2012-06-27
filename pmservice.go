package libclient3

import "os"

// this struct pointer implements Service.
type PMService struct {
	connection *Connection
}

// connects to the ProcessManager service.
func ConnectPM() (pm *PMService, err error) {
	// connect to ProcessManager.
	conn, err := Connect("/system/socket/ProcessSocket")
	if err == nil {
		pm = &PMService{conn}
		conn.service = pm
	}
	return
}

// creates, connects, registers, and loops
func RunPM(data map[string]interface{}) error {
	pm, err := ConnectPM()
	if err != nil {
		return err
	}
	pm.Register(data)
	pm.Loop()
	return nil
}

// register to ProcessManager.
func (pm *PMService) Register(data map[string]interface{}) {
	flexibleData := make(map[string]interface{}, len(data)+1)
	for key, val := range data {

		// only allow strings.
		switch val.(type) {
		case string:
		default:
			continue
		}
		flexibleData[key] = val
	}
	flexibleData["pid"] = os.Getpid()
	pm.connection.Send("register", flexibleData)
}

// loop.
func (pm *PMService) Loop() {
	pm.connection.Run()
	os.Exit(1)
}

// handle a JSON event.
func (pm *PMService) handleEvent(command string, params map[string]interface{}) {
	if PMEvents[command] != nil {
		PMEvents[command](pm, params)
	}
}
