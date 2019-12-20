package mqttparser

import (
	"log"
)

//Connect returns the content of the connect message
type Connect struct {
	Command       string
	WillRetain    bool
	Qos           uint
	CleanStart    bool
	KeepAlive     uint16
	SessionExp    uint32
	ReceiveMax    uint16
	MaxPacketSize uint32
	TopicAliasMax uint16
	ReqResInfo    int
	ReqProbInfo   int
	Username      string
	ClientID      string
	Password      string
}

//HandleConnect handles the connect message
func HandleConnect(b []byte) (Connect, error) {

	log.Printf("%s", b)
	c := Connect{Command: "Connect"}

	//check variable header for correctness
	if b[0] != 0x00 || b[1] != 0x04 || b[2] != 0x4d || b[3] != 0x51 || b[4] != 0x54 || b[5] != 0x54 {
		log.Println("error in CONNECT")
	}

	ver := b[6]
	if ver != 5 {
		log.Println("error < 5 not supported")
	}

	//Check UserName Flag
	un := false
	if b[7]&0x80 == 0x80 {
		un = true
	}

	log.Printf("username %t\n", un)

	//Check Password flag
	pwd := false
	if b[7]&0x40 == 0x40 {
		pwd = true
	}

	log.Printf("pasword %t\n", pwd)

	//Check Will Retain flag
	c.WillRetain = false
	if b[7]&0x20 == 0x20 {
		c.WillRetain = true
	}
	log.Printf("willRetain %t\n", c.WillRetain)

	//Check will QOS
	c.Qos = uint(b[7]&0x10 + b[7]&0x08)
	log.Printf("will qos %d", c.Qos)

	//Check will flag
	willFlag := false
	if b[7]&0x04 == 0x04 {
		willFlag = true
	}

	log.Printf("will %t\n", willFlag)

	//Check clean start
	c.CleanStart = false
	if b[7]&0x02 == 0x02 {
		c.CleanStart = true
	}

	log.Printf("clean start %t\n", c.CleanStart)

	if b[7]&0x01 == 0x01 {
		log.Println("ERROR")
	}

	c.KeepAlive = getUint16(b[8], b[9])
	log.Printf("%d\n", c.KeepAlive)

	propLen := b[10]
	log.Printf("properties len %d", propLen)

	//Scan properties
	m := 11
	for i := m; i < int(propLen)+m; i++ {
		bi := b[i]

		switch int(bi) {
		//Session Expiry Interval
		case 0x11:
			c.SessionExp = getUint32(b[i+1], b[i+2], b[i+3], b[i+4])
			i += 4
			log.Printf("session expires %d", c.SessionExp)

		case 0x21:
			//Receive Maximum
			c.ReceiveMax = getUint16(b[i+1], b[i+2])
			i += 2
			log.Printf("receive maximum %d", c.ReceiveMax)

		case 0x27:
			//Maximum Packet Size
			c.MaxPacketSize = getUint32(b[i+1], b[i+2], b[i+3], b[i+4])
			i += 4
			log.Printf("maximum packet size %d", c.MaxPacketSize)

		case 0x22:
			//Topic Alias Maximum
			c.TopicAliasMax = getUint16(b[i+1], b[i+2])
			i += 4
			log.Printf("receive maximum %d", c.TopicAliasMax)

		case 0x19:
			//Request response info
			c.ReqResInfo = int(b[i+1])
			i++
			log.Printf("Request response information %d", c.ReqResInfo)

		case 0x17:
			//Request Problem Information
			c.ReqProbInfo = int(b[i+1])
			i++
			log.Printf("Request Problem Information %d", c.ReqProbInfo)

		case 0x26:
			unBuf := make([]byte, 2)
			unBuf[0] = b[i+1]
			unBuf[1] = b[i+2]
			c.Username = string(unBuf)

			log.Printf("Username %s", c.Username)
		case 0x15:
			//Authentication Extension
			log.Printf("Authentication Extension")
		case 0x16:
			//Authentication Data
			log.Printf("Authentication Data")
		default:
			log.Println("ERROR")

		}

		m = i
	}

	clientID, a := getUtf8(b[m:])
	c.ClientID = clientID
	m += a
	log.Printf("ClientID %s %d\n", c.ClientID, a)

	if willFlag == true {

	}

	if un == true {
		username, a := getUtf8(b[m:])
		c.Username = username
		log.Printf("username %s\n", username)
		m += a
	}

	if pwd == true {
		password, _ := getUtf8(b[m:])
		c.Password = password
		log.Printf("password %s\n", password)
	}

	return c, nil
}
