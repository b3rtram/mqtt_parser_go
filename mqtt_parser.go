package mqttparser

import (
	"encoding/binary"
	"errors"
	"fmt"
)

const (
	// Time to wait before starting closing clients when in LD mode.
	connect     = 0x10
	publish     = 0x30
	subscribe   = 0x82
	suback      = 0x09
	unsubscribe = 0xA0
	disconnect  = 0xe0
	pingreq     = 0xc0
)

//MqttCommand describes
type MqttCommand struct {
	Command string
	MqttLen int
}

//GetCommand take a look into the bytearray
func GetCommand(b []byte) (MqttCommand, int, error) {

	mqttLen, read := getVarByteInt(b[1:])
	fmt.Printf("%d\n", mqttLen)

	pos := read + 1

	switch b[0] {
	case connect:
		return MqttCommand{Command: "Connect", MqttLen: mqttLen}, pos, nil

	case publish:
		return MqttCommand{Command: "Publish", MqttLen: mqttLen}, pos, nil

	case subscribe:
		return MqttCommand{Command: "Subscribe", MqttLen: mqttLen}, pos, nil

	case unsubscribe:
		fmt.Printf("test\n")

	case pingreq:
		return MqttCommand{Command: "Pingreq", MqttLen: mqttLen}, pos, nil
		fmt.Printf("pingreq\n")

	case disconnect:
		
		fmt.Printf("disconnect\n")
	}

	return MqttCommand{Command: ""}, 0, errors.New("This parser does not support this kind of message")
}

func getVarByteInt(bs []byte) (int, int) {
	multiplier := 1
	value := 0
	a := 0
	for {
		encodedByte := bs[a]
		value += (int(encodedByte) & 127) * multiplier
		if multiplier > 128*128*128 {
			break
		}

		multiplier *= 128
		a++
		if (encodedByte & 128) == 0 {
			break
		}

	}

	return value, a
}

func getUtf8(bs []byte) (string, int) {
	len := getUint16(bs[0], bs[1])

	clientID := make([]byte, len)
	a := 0
	for i := 0; i < int(len); i++ {
		clientID[i] = bs[i+2]
		a++
	}

	return string(clientID), a + 2
}

func getUint16(b1 byte, b2 byte) uint16 {

	bs := make([]byte, 2)
	bs[0] = b1
	bs[1] = b2

	return binary.BigEndian.Uint16(bs)
}

func getUint32(b1 byte, b2 byte, b3 byte, b4 byte) uint32 {

	bs := make([]byte, 4)
	bs[0] = b1
	bs[1] = b2
	bs[2] = b3
	bs[3] = b4

	return binary.BigEndian.Uint32(bs)
}
