package mqttparser

import (
	"fmt"
	"log"
)

//Publish test
type Publish struct {
	Command string
	Topic   string
	Message []byte
}

//HandlePublish is cool
func HandlePublish(b []byte, mqttLen int) (Publish, error) {
	p := Publish{Command: "Publish"}

	topic, len := getUtf8(b)
	p.Topic = topic
	log.Printf("topic: %s %d\n", topic, len)

	pos := len + 2
	payLen := mqttLen - pos
	msg := b[pos : pos+payLen]
	p.Message = msg
	fmt.Printf("pub payload: %s\n", string(msg))

	return p, nil
}
