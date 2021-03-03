package rabbitmq

import (
	"testing"
)

const address = "amqp://admin:yunjikeji@10.0.3.168:5672//mcollective"

func TestRoutingKey(t *testing.T) {
	c := NewClient(address)

	//c.RoutingKey("mcollective_broadcast", "discovery", "")
	//c.RoutingKey("mcollective_broadcast", "rpcutil", "")
	c.RoutingKey("mcollective_broadcast", "registration", "")

	defer c.Close()
}
