package RabbitCallback

import "github.com/streadway/amqp"



type Consume interface {
	Consume(ch *amqp.Channel, QueueName string)
	Callback(data interface{})
}

