package RabbitCallback

import (
	"log"
)


func AddressCallBack(data interface{}) error  {
	log.Println("address回调收到消息 ",data)
	return nil
}

func UserCallBack(data interface{}) error  {
	log.Println("user回调收到消息 ",data)
	return nil
}

func ZoneCallBack(data interface{}) error  {
	log.Println("zone回调收到消息 ",data)
	return nil
}