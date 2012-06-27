package libclient3

import "os"

// this struct pointer implements Service.
type LMService struct {
	connection *Connection
}

// connects to the LaunchManager service.
func ConnectLM() (lm *LMService, err error) {

	// setup LaunchManager event handlers.
	if LMEvents == nil {
		CreateLMEvents()
	}

	// connect to LaunchManager.
	conn, err := Connect("/system/socket/LaunchSocket")
	if err == nil {
		lm = &LMService{conn}
		conn.service = lm
	}
	return
}

// creates, connects, registers, and loops
func RunLM() error {
	lm, err := ConnectLM()
	if err != nil {
		return err
	}
	lm.Register(nil)
	lm.Loop()
	return nil
}

// register to LaunchManager.
func (lm *LMService) Register(_ map[string]interface{}) {
	lm.connection.Send("register", map[string]interface{}{
		"pid": os.Getpid(),
	})
}

// loop.
func (lm *LMService) Loop() {
	lm.connection.Run()
	os.Exit(1)
}

// handle a JSON event.
func (*LMService) handleEvent(command string, params map[string]interface{}) {
	if LMEvents[command] != nil {
		LMEvents[command](params)
	}
}
