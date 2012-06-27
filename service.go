package libclient3

type Service interface {

	// handle a JSON event.
	handleEvent(command string, params map[string]interface{})
}
