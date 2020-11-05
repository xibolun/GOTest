package basic

import (
	"fmt"
	"net/rpc"
	"testing"
)

func TestCustomer(t *testing.T) {
	client, err := rpc.Dial("tcp4", ":1234")

	if err != nil {
		t.Error(err)
	}

	var replay string
	if err := client.Call("HelloService.Hello", "world", &replay); err != nil {
		t.Error(err)
	}

	fmt.Println(replay)
}
