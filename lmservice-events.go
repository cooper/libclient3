package libclient3

var LMEvents map[string]func(map[string]interface{})

func CreateLMEvents() {
	LMEvents = make(map[string]func(map[string]interface{}))
}
