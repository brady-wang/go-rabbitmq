package Error

import (
	"fmt"
	"log"
	"os"
)

func init() {
	logFile, err := os.OpenFile("./log/rabbitmq.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("open log file failed, err:", err)
		return
	}
	log.SetOutput(logFile)
	log.SetFlags(log.Lshortfile | log.Ldate | log.Ltime)
	//log.SetPrefix("rabbitmq")
}

func FailOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}
