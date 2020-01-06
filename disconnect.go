package mqttparser

//Disconnect returns the content of the connect message
type Disconnect struct {
	Command string
}

//HandleDisconnect handles the connect message
func HandleDisconnect(b []byte) (Disconnect, error) {
	return Disconnect{Command: "Disconnect"}, nil
}
