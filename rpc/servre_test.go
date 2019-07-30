package rpc

import (
	"log"
	"net"
	"net/rpc"
	"testing"
)

func TestServerStart(t *testing.T) {
	_ = rpc.RegisterName("HelloService", new(HelloService))

	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal(err)
	}

	conn, err := listener.Accept()
	if err != nil {
		log.Fatal(err)
	}

	rpc.ServeConn(conn)
}
