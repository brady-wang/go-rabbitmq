package Topic

import (
	"github.com/streadway/amqp"
	"log"
	"mq/Error"
)

type Consumer struct {
	HandleFunc  func (data interface{}) error
}

func (c *Consumer) Callback(data interface{}) {
	log.Printf("接受到数据 %v",data)
	handleFunc := c.HandleFunc
	_ = handleFunc(data)
}

func (c *Consumer) Consume(ch *amqp.Channel, QueueName string) {
	msgs, err := ch.Consume(
		QueueName, // queue
		"",        // consumer
		false,     // auto ack
		false,     // exclusive
		false,     // no local
		false,     // no wait
		nil,       // args
	)
	Error.FailOnError(err, "Failed to register a consumer")

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			if true {
				d.Ack(true)
				log.Printf(" [%s]收到消息 %s", QueueName, d.Body)
				c.Callback(string(d.Body))
			} else {
				d.Ack(true)
				log.Printf(" [%s]mq异常 %s", QueueName, d.Body)
			}

		}
	}()

	log.Printf(" [%s] Waiting for logs. To exit press CTRL+C", QueueName)
	<-forever
}
