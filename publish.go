package mqttparser

import (
	"fmt"
	"log"
)

//Publish test
type Publish struct {
	Command     string
	Topic       string
	Message     []byte
	CompleteMsg []byte
}

//HandlePublish is cool
func HandlePublish(completeB []byte, header int, mqttLen int) (Publish, error) {

	p := Publish{Command: "Publish"}
	b := completeB[header:]

	topic, len := getUtf8(b)
	p.Topic = topic
	log.Printf("topic: %s %d\n", topic, len)

	pos := len + 1
	payLen := mqttLen - pos
	msg := b[pos : pos+payLen]
	p.Message = msg
	p.CompleteMsg = completeB[:mqttLen+2]
	fmt.Printf("pub payload: %s\n", string(msg))

	return p, nil
}
