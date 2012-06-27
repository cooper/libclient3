package libclient3

type Service interface {

	// handle a JSON event.
	handleEvent(command string, params map[string]interface{})

	// register the process to the service.
	Register(params map[string]interface{})

	// run the connection loop.
	Loop()
}
