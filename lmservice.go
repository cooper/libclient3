package libclient3

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

func (*LMService) handleEvent(command string, params map[string]interface{}) {
	if LMEvents[command] != nil {
		LMEvents[command](params)
	}
}
