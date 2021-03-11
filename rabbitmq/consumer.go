package rabbitmq

import (
	"log"
	"time"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func (c *RClient) Consumer(exchangeName, routeKey, queueName string) {
	//_, err := c.ch.QueueDeclare(queueName, true, false, true, false, nil)
	//failOnError(err, "Failed to declare a queue")

	if exchangeName != "" && routeKey != "" {
		log.Printf("Binding queue %s to exchange %s with routing key %s", queueName, exchangeName, routeKey)
		err := c.ch.QueueBind(queueName, routeKey, exchangeName, false, nil)
		failOnError(err, "Failed to bind a queue")
	}

	msgs, err := c.ch.Consume(queueName, "", true, false, false, false, nil)
	failOnError(err, "Failed to register a consumer")

	for {
		select {
		case d := <-msgs:
			result := McoResult{}
			//log.Printf("get mq message from %s, message: %s", queueName, string(d.Body))
			if err = UnRubyMarshal(d.Body, &result); err != nil {
				log.Printf("ruby unmarshal fail, %s", err.Error())
				continue
			}

			log.Printf("MsgTime is %s\n", time.Unix(result.MsgTime, 0).Format("2006-01-02 15:04:05"))
			log.Printf("RequestID is %s\n", result.RequestID)
			log.Printf("SenderAgent is %s\n", result.SenderAgent)
			log.Printf("SenderID is %s\n", result.SenderID)
			log.Printf("----------------> done")
			//log.Printf("Hash is %s\n", result.Hash)
			//log.Printf("TTL is %d\n", result.TTL)
			//log.Printf("Agent is %s\n", result.Agent)
			//log.Printf("Collective is %s\n", result.Collective)
			//log.Printf("CallerID is %s\n", result.CallerID)
			//log.Printf("Compound is %v\n", result.Compound)
			//log.Printf("Fact is %v\n", result.Fact)
			//log.Printf("CfClass is %v\n", result.CfClass)
			//log.Printf("Body is %s\n", string(result.Body))
		}
	}
}

func (c *RClient) RoutingKeyConsumer(exchangeName, routeKey, queueName string) {
	q, err := c.ch.QueueDeclare(queueName, false, false, true, false, nil)
	failOnError(err, "Failed to declare a queue")

	log.Printf("Binding queue %s to exchange %s with routing key %s", queueName, exchangeName, routeKey)

	err = c.ch.QueueBind(q.Name, routeKey, exchangeName, false, nil)
	failOnError(err, "Failed to bind a queue")

	msgs, err := c.ch.Consume(q.Name, "", true, false, false, false, nil)
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
			log.Printf("MsgTime is %s\n", time.Unix(result.MsgTime, 0).Format("2006-01-02 15:04:05"))
			log.Printf("RequestID is %s\n", result.RequestID)
			log.Printf("SenderAgent is %s\n", result.SenderAgent)
			log.Printf("SenderID is %s\n", result.SenderID)
			log.Printf("Hash is %s\n", result.Hash)
			log.Printf("TTL is %d\n", result.TTL)
			log.Printf("Agent is %s\n", result.Agent)
			log.Printf("Collective is %s\n", result.Collective)
			log.Printf("CallerID is %s\n", result.CallerID)
			log.Printf("Compound is %v\n", result.Compound)
			log.Printf("Fact is %v\n", result.Fact)
			log.Printf("CfClass is %v\n", result.CfClass)
			log.Printf("----------------> done")

			//body := make([]string, 0)
			//if err = UnRubyMarshal([]byte(result.Body), &body); err != nil {
			//	log.Printf("ruby unmarshal fail, %s", err.Error())
			//	continue
			//}
			//log.Printf("Body is %v\n", body)

		}
	}()

	log.Printf(" [*] Waiting for logs. To exit press CTRL+C")
	<-forever
}
