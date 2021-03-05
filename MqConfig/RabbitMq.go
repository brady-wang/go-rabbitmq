package MqConfig

var MqSource string = "amqp://guest:guest@192.168.1.99:5672/"

type RabbitMqConfig struct {
	ExchangeName string
	QueueName    string
	RouteName    string
}

var UserMq = &RabbitMqConfig{
	ExchangeName: "exchange_user",
	QueueName:    "queue_user",
	RouteName:    "route_user",
}

var AddressMq = &RabbitMqConfig{
	ExchangeName: "exchange_address",
	QueueName:    "queue_address",
	RouteName:    "route_address",
}

var ZoneMq = &RabbitMqConfig{
	ExchangeName: "exchange_zone",
	QueueName:    "queue_zone",
	RouteName:    "route_zone",
}
