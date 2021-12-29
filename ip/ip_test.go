package ip

import (
	"net"
	"testing"
)

func TestIP(t *testing.T) {
	str := "192.168.0.1"

	ip := net.ParseIP(str)

	t.Log(ip.DefaultMask().String())
}
