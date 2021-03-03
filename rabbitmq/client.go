package rabbitmq

import "github.com/streadway/amqp"

type RClient struct {
	c  *amqp.Connection
	ch *amqp.Channel
}

// NewClient
func NewClient(address string) *RClient {
	conn, err := amqp.Dial(address)
	if err != nil {
		failOnError(err, "Failed to connect to RabbitMQ")
		return nil
	}

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")

	return &RClient{
		c:  conn,
		ch: ch,
	}
}

// Close
func (c *RClient) Close() {
	_ = c.c.Close()
	_ = c.ch.Close()
}
