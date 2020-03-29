package pubsub

import (
	"github.com/geiqin/supports/helper"
	"github.com/micro/go-micro/broker"
	"log"
)

var pubSub broker.Broker

type EventType string

//消息注册
func Register(myBroker broker.Broker) {
	pubSub = myBroker
	if err := pubSub.Connect(); err != nil {
		log.Println("register broker connect error: %v\n", err)
	}
}

//消息发布
func Publish(eventName EventType, storeId int64, data string, headers ...map[string]string) error {
	heads := make(map[string]string)
	if storeId > 0 {
		heads["store_id"] = helper.Int64ToString(storeId)
	}
	if headers != nil {
		for k, v := range headers[0] {
			if k != "store_id" {
				heads[k] = v
			}
		}
	}
	msg := &broker.Message{
		Header: heads,
		Body:   []byte(data),
	}
	err := pubSub.Publish(string(eventName), msg)
	return err
}

//订阅消息
func Subscribe(eventName EventType, handler broker.Handler) (broker.Subscriber, error) {
	sub, err := pubSub.Subscribe(string(eventName), handler)
	return sub, err
}
