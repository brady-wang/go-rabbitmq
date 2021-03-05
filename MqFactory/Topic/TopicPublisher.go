package Topic

import (
	"github.com/streadway/amqp"
	"log"
	"mq/Error"
)


func Publish(body string ,ch*amqp.Channel,exchangeName string,queueName string, routeName string)  {
	err := ch.Publish(
		exchangeName,          // exchange
		routeName, // routing key
		false, // mandatory
		false, // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	Error.FailOnError(err, "Failed to publish a message")

	log.Printf(" [%s] Sent %s",queueName, body)
}
