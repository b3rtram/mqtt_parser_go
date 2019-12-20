package mqttparser

import "fmt"

//Subscribe test
type Subscribe struct {
	Command  string
	PacketID uint16
	SubID    int
	Topic    []string
}

//HandleSubscribe is cool
func HandleSubscribe(b []byte, mqttLen int) (Subscribe, error) {
	s := Subscribe{Command: "Subscribe"}
	s.Topic = make([]string, 1)
	pos := 0
	s.PacketID = getUint16(b[pos], b[pos+1])
	fmt.Printf("packetID %d\n", s.PacketID)
	pos += 2
	propLen := b[pos]
	pos++
	s.SubID = 0

	for i := 0; i < int(propLen-1); i++ {
		bi := b[pos+i]
		pos++
		switch int(bi) {
		case 0x0b:
			var r int
			s.SubID, r = getVarByteInt(b[pos:])
			fmt.Printf("subID %d %d\n", s.SubID, r)
			pos += r
		case 0x26:
			user, r := getUtf8(b)
			fmt.Printf("user: %s\n", user)
			pos += r
		}
	}

	for {
		topic, r := getUtf8(b[pos:])
		s.Topic = append(s.Topic, topic)
		fmt.Printf("topic: %s %d\n", topic, r)
		pos += r + 1

		if pos > mqttLen {
			break
		}
	}

	return s, nil

}
