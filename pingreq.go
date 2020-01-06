package mqttparser

//PingReq is a
type PingReq struct {
	Command string
}

//HandlePingReq handles
func HandlePingReq(b []byte) (PingReq, error) {
	return PingReq{Command: "PingReq"}, nil
}
