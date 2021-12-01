package snmp

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"path/filepath"
	"testing"

	g "github.com/gosnmp/gosnmp"
)

func TestSNMPReceiver(t *testing.T) {
	for {
		Receive()
	}
}

func Receive() {
	flag.Usage = func() {
		fmt.Printf("Usage:\n")
		fmt.Printf("   %s\n", filepath.Base(os.Args[0]))
		flag.PrintDefaults()
	}

	param := g.Default
	//param.Community = "cloudboot"

	tl := g.NewTrapListener()
	tl.OnNewTrap = myTrapHandler
	tl.Params = param
	tl.Params.Logger = log.New(os.Stdout, "", 0)

	err := tl.Listen("10.0.2.1:161")
	if err != nil {
		log.Panicf("error in listen: %s", err)
	}
}

func myTrapHandler(packet *g.SnmpPacket, addr *net.UDPAddr) {
	log.Printf("got trapdata from %s\n", addr.IP)
	for _, v := range packet.Variables {
		switch v.Type {
		case g.OctetString:
			b := v.Value.([]byte)
			fmt.Printf("OID: %s, string: %x\n", v.Name, b)

		default:
			log.Printf("trap: %+v\n", v)
		}
	}
}
