package libclient3

var PMEvents map[string]func(map[string]interface{})

func CreatePMEvents() {
	PMEvents = make(map[string]func(map[string]interface{}))
}
