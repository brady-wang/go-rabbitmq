package MqService

import (
	"github.com/streadway/amqp"
	"mq/RabbitCallback"
	"mq/MqConfig"
	"mq/MqFactory/Topic"
	"reflect"
	"sync"
	"time"
)

type MqSyncMethod struct {


}



func Sync(funcName string, ch *amqp.Channel, wg *sync.WaitGroup) () {

	value := reflect.ValueOf(&MqSyncMethod{})
	methodValue := value.MethodByName(funcName)
	args := make([]reflect.Value,0)
	args = append(args, reflect.ValueOf(ch),reflect.ValueOf(wg) )
	methodValue.Call(args)
}



func (m *MqSyncMethod) SyncAddress(ch *amqp.Channel, wg *sync.WaitGroup) {
	defer wg.Done()
	Topic.ExchangeDeclare(ch, MqConfig.AddressMq.ExchangeName) // 声明交换机
	Topic.QueueDeclare(ch, MqConfig.AddressMq.QueueName)
	Topic.BindRoute(ch, MqConfig.AddressMq.ExchangeName, MqConfig.AddressMq.QueueName, MqConfig.AddressMq.RouteName)
	consumer := Topic.Consumer{HandleFunc: RabbitCallback.AddressCallBack}
	consumer.Consume(ch, MqConfig.AddressMq.QueueName)

}

func (m *MqSyncMethod) SyncUser(ch *amqp.Channel, wg *sync.WaitGroup) {
	defer wg.Done()
	Topic.ExchangeDeclare(ch, MqConfig.UserMq.ExchangeName) // 声明交换机
	Topic.QueueDeclare(ch, MqConfig.UserMq.QueueName)
	Topic.BindRoute(ch, MqConfig.UserMq.ExchangeName, MqConfig.UserMq.QueueName, MqConfig.UserMq.RouteName)
	consumer := Topic.Consumer{HandleFunc: RabbitCallback.UserCallBack}
	consumer.Consume(ch, MqConfig.UserMq.QueueName)
}

func (m *MqSyncMethod) SyncZone(ch *amqp.Channel, wg *sync.WaitGroup) {
	defer wg.Done()
	Topic.ExchangeDeclare(ch, MqConfig.ZoneMq.ExchangeName) // 声明交换机
	Topic.QueueDeclare(ch, MqConfig.ZoneMq.QueueName)
	Topic.BindRoute(ch, MqConfig.ZoneMq.ExchangeName, MqConfig.ZoneMq.QueueName, MqConfig.ZoneMq.RouteName)
	//Topic.Consume(ch, MqConfig.ZoneMq.QueueName)
	consumer := Topic.Consumer{HandleFunc: RabbitCallback.ZoneCallBack}
	consumer.Consume(ch, MqConfig.ZoneMq.QueueName)
}




func (m *MqSyncMethod) PushZone(message string,ch *amqp.Channel,wg *sync.WaitGroup)  {
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