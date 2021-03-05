package MqService

import (
	"github.com/streadway/amqp"
	"mq/MqConfig"
	"mq/MqFactory/Topic"
	"sync"
	"time"
)

func Sync(funcName string, ch *amqp.Channel, wg *sync.WaitGroup) () {

	switch funcName {
	case "SyncZone":
		SyncZone(ch, wg)

	case "SyncAddress":
		SyncAddress(ch, wg)

	case "SyncUser":
		SyncUser(ch, wg)

	default:
		SyncZone(ch, wg)

	}

}

func SyncAddress(ch *amqp.Channel, wg *sync.WaitGroup) {
	defer wg.Done()
	Topic.ExchangeDeclare(ch, MqConfig.AddressMq.ExchangeName) // 声明交换机
	Topic.QueueDeclare(ch, MqConfig.AddressMq.QueueName)
	Topic.BindRoute(ch, MqConfig.AddressMq.ExchangeName, MqConfig.AddressMq.QueueName, MqConfig.AddressMq.RouteName)
	Topic.Consume(ch, MqConfig.AddressMq.QueueName)
}

func SyncUser(ch *amqp.Channel, wg *sync.WaitGroup) {
	defer wg.Done()
	Topic.ExchangeDeclare(ch, MqConfig.UserMq.ExchangeName) // 声明交换机
	Topic.QueueDeclare(ch, MqConfig.UserMq.QueueName)
	Topic.BindRoute(ch, MqConfig.UserMq.ExchangeName, MqConfig.UserMq.QueueName, MqConfig.UserMq.RouteName)
	Topic.Consume(ch, MqConfig.UserMq.QueueName)
}

func SyncZone(ch *amqp.Channel, wg *sync.WaitGroup) {
	defer wg.Done()
	Topic.ExchangeDeclare(ch, MqConfig.ZoneMq.ExchangeName) // 声明交换机
	Topic.QueueDeclare(ch, MqConfig.ZoneMq.QueueName)
	Topic.BindRoute(ch, MqConfig.ZoneMq.ExchangeName, MqConfig.ZoneMq.QueueName, MqConfig.ZoneMq.RouteName)
	Topic.Consume(ch, MqConfig.ZoneMq.QueueName)
}

func PushZone(message string,ch *amqp.Channel,wg *sync.WaitGroup)  {
	defer wg.Done()
	Topic.ExchangeDeclare(ch, MqConfig.ZoneMq.ExchangeName) // 声明交换机
	Topic.QueueDeclare(ch, MqConfig.ZoneMq.QueueName)
	Topic.BindRoute(ch, MqConfig.ZoneMq.ExchangeName, MqConfig.ZoneMq.QueueName, MqConfig.ZoneMq.RouteName)

	ticker := time.NewTicker(time.Second * 1) // 运行时长
	ch1 := make(chan int)
	go func() {
		var x int
		for x < 1000 {
			select {
			case <-ticker.C:
				x++
				Topic.Publish(message,ch,MqConfig.ZoneMq.ExchangeName,MqConfig.ZoneMq.QueueName,MqConfig.ZoneMq.RouteName)
			}
		}
		ticker.Stop()
		ch1 <- 0
	}()
	<-ch1

}