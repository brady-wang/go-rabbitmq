package Topic

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
	"mq/Error"
)

func GetMqConnection(mqSource string) *amqp.Connection {
	conn, err := amqp.Dial(mqSource)
	Error.FailOnError(err, "Failed to connect to RabbitMQ")
	return conn
}

func GetMqChannel(conn *amqp.Connection) *amqp.Channel {
	ch, err := conn.Channel()
	Error.FailOnError(err, "Failed to open a channel")
	return ch
}

func ExchangeDeclare(ch *amqp.Channel, ExchangeName string) {
	err := ch.ExchangeDeclare(
		ExchangeName, // name
		"topic",      // type
		true,         // durable
		false,        // auto-deleted
		false,        // internal
		false,        // no-wait
		nil,          // arguments
	)
	Error.FailOnError(err, "Failed to declare an exchange")
}
func QueueDeclare(ch *amqp.Channel, QueueName string) {
	_, err := ch.QueueDeclare(
		QueueName, // name
		true,      // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	Error.FailOnError(err, "Failed to declare a queue")
}
func BindRoute(ch *amqp.Channel, ExchangeName string, QueueName string, RouteName string) {
	//for _, s := range RouteList {
	fmt.Println(ExchangeName,QueueName,RouteName)
	log.Printf("Binding queue %s to exchange %s with routing key %s", QueueName, ExchangeName, RouteName)
	err := ch.QueueBind(
		QueueName,    // queue name
		RouteName,    // routing key
		ExchangeName, // exchange
		false,
		nil)
	Error.FailOnError(err, "Failed to bind a queue")
	//}
}
