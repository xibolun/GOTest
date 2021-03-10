package rabbitmq

import (
	"testing"

	"xibolun/gotest/basic"
)

const McoVirtualAddress = "amqp://admin:yunjikeji@10.0.3.168:5672//mcollective"
const BizVirtualAddress = "amqp://admin:yunjikeji@10.0.3.175:5672//mcollective"
const ReplyQueue = "SEVEN_TEST_2"

func TestRegistration(t *testing.T) {
	c := NewClient(McoVirtualAddress)

	c.RoutingKeyConsumer("mcollective_broadcast", "registration", "")

	defer c.Close()
}

func TestDiscovery(t *testing.T) {
	c := NewClient(McoVirtualAddress)

	c.RoutingKeyConsumer("mcollective_broadcast", "discovery", "")

	defer c.Close()
}

func TestShell(t *testing.T) {
	c := NewClient(McoVirtualAddress)

	c.RoutingKeyConsumer("mcollective_broadcast", "shell", "")

	defer c.Close()
}

func TestRpcUtil(t *testing.T) {
	c := NewClient(McoVirtualAddress)

	c.RoutingKeyConsumer("mcollective_broadcast", "rpcutil", "")

	defer c.Close()
}

func TestDirectChangeConsumer(t *testing.T) {
	c := NewClient(McoVirtualAddress)

	c.RoutingKeyConsumer("mcollective_directed", "10.0.3.168", "")

	defer c.Close()
}

func TestConsumer(t *testing.T) {
	c := NewClient(BizVirtualAddress)

	res := make(chan interface{})

	c.Consumer("", "", ReplyQueue, res)

	value := <-res
	t.Log(basic.ToJsonString(value.(McoResult)))

	defer c.Close()
}
