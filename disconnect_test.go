package mqttparser

import (
	"fmt"
	"testing"
)

func TestHandleDisConnect(t *testing.T) {

	b := [...]byte{0xe0, 0x00}
	c, err := HandleDisconnect(b[2:])

	if err != nil {
		t.Error("Faiure")
	}

	if c.Command != "Disconnect" {
		t.Error("Disconnect failure")
	}

	fmt.Printf("Disconnect\n")
}
