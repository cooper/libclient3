package libclient3

import (
	"bufio"
	"encoding/json"
	"net"
	"os"
)

type Connection struct {
	path     string
	socket   net.Conn
	incoming *bufio.Reader
	service  Service
}

func Connect(path string) (conn *Connection, err error) {

	// check if file exists. if not, bail.
	_, err = os.Lstat(path)
	if err != nil {
		return
	}

	// resolve the address
	addr, err := net.ResolveUnixAddr("unix", path)
	if err != nil {
		return
	}

	// connect
	unixConn, err := net.DialUnix("unix", nil, addr)
	if err != nil {
		return
	}

	conn = &Connection{
		path:     path,
		socket:   unixConn,
		incoming: bufio.NewReader(unixConn),
	}

	return
}

// send a JSON event
func (conn *Connection) Send(command string, params map[string]interface{}) bool {
	b, err := json.Marshal([]interface{}{command, params})
	if err != nil {
		return false
	}
	b = append(b, '\n')
	if _, err = conn.socket.Write(b); err != nil {
		return false
	}
	return true
}

// read data from a connection continuously
func (conn *Connection) Run() {
	for {
		line, _, err := conn.incoming.ReadLine()
		if err != nil {
			return
		}
		conn.handleEvent(line)
	}
}

// handle a JSON event
func (conn *Connection) handleEvent(data []byte) bool {
	var i interface{}
	err := json.Unmarshal(data, &i)
	if err != nil {
		return false
	}

	// should be an array.
	c := i.([]interface{})

	command := c[0].(string)
	params := c[1].(map[string]interface{})

	// if a handler for this command exists, run it
	conn.service.handleEvent(command, params)

	return true
}
