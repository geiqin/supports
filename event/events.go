package event

import (
	"github.com/micro/go-micro/broker"
	"log"
	"sync"
)

var pubSub broker.Broker
var once sync.Once

type EventType string

//消息注册
func Register(myBroker broker.Broker) {
	once.Do(func() {
		pubSub = myBroker
		if err := pubSub.Connect(); err != nil {
			log.Println("register broker connect error: %v\n", err)
		}
	})
}

//消息发布
func Publish(eventName EventType, data string) error {
	msg := &broker.Message{
		Body: []byte(data),
	}
	err := pubSub.Publish(string(eventName), msg)
	return err
}

//订阅消息
func Subscribe(eventName EventType, handler broker.Handler) (broker.Subscriber, error) {
	sub, err := pubSub.Subscribe(string(eventName), handler)
	return sub, err
}
