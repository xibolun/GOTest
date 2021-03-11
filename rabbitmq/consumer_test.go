package rabbitmq

import (
	"testing"
)

const BizVirtualAddress = "amqp://admin:yunjikeji@10.0.3.175:5672//mcollective"
const ReplyQueue = "SEVEN_TEST_3"

func TestRegistration(t *testing.T) {
	c := NewClient(BizVirtualAddress)

	c.RoutingKeyConsumer("mcollective_broadcast", "registration", "")

	defer c.Close()
}

func TestDirectChangeConsumer(t *testing.T) {
	c := NewClient(BizVirtualAddress)

	c.RoutingKeyConsumer("mcollective_directed", "10.0.3.168", "")

	defer c.Close()
}

func TestConsumer(t *testing.T) {
	c := NewClient(BizVirtualAddress)

	forever := make(chan bool)

	c.Consumer("", "", ReplyQueue)

	<-forever

}
