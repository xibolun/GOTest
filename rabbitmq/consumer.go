package rabbitmq

import (
	"log"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func (c *RClient) RoutingKey(exchangeName, routeKey, queueName string) {

	q, err := c.ch.QueueDeclare(
		queueName, // name
		false,     // durable
		false,     // delete when unused
		true,      // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	failOnError(err, "Failed to declare a queue")

	log.Printf("Binding queue %s to exchange %s with routing key %s", queueName, exchangeName, routeKey)

	err = c.ch.QueueBind(
		q.Name,       // queue name
		routeKey,     // routing key
		exchangeName, // exchange
		false,
		nil)
	failOnError(err, "Failed to bind a queue")

	msgs, err := c.ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto ack
		false,  // exclusive
		false,  // no local
		false,  // no wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			result := McoResult{}
			if err = UnRubyMarshal(d.Body, &result); err != nil {
				log.Printf("ruby unmarshal fail, %s", err.Error())
				continue
			}
			log.Printf("result message is %s ", string(d.Body))
			log.Printf("unmarshl value is %v", result)
		}
	}()

	log.Printf(" [*] Waiting for logs. To exit press CTRL+C")
	<-forever
}
