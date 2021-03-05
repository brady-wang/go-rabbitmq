package main

import (
	"github.com/streadway/amqp"
	"log"
	"mq/MqConfig"
	"mq/MqFactory/Topic"
	"mq/MqService"
	"sync"
)

var (
	conn *amqp.Connection
	ch   *amqp.Channel
	wg   sync.WaitGroup
)

func init() {
	conn = Topic.GetMqConnection(MqConfig.MqSource)
	ch = Topic.GetMqChannel(conn)
}

func main() {



	//wg.Add(1)
	//go MqService.PushZone("hello world",ch,&wg)

	var mqList =  make([]string,0)
	mqList = append(mqList, "SyncZone","SyncUser","SyncAddress")

	for _,v := range mqList{
		wg.Add(1)
		log.Printf("[%s] 消费开启\n", v)
		go MqService.Sync(v,ch,&wg)
	}

	wg.Wait()
	defer conn.Close()
	defer ch.Close()
}
