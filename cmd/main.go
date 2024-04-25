package main

import (
	"fmt"

	"github.com/rafaelsouzaribeiro/consumer/pkg"
	exec "github.com/rafaelsouzaribeiro/consumer/pkg/kafka"
)

func main() {
	consumer := exec.NewReader([]string{"springboot:9092"})
	msg := pkg.ReadMessage{
		Topic:   "contact-adm-insert",
		GroupID: "contact",
	}
	go consumer.Receive(&msg, handleMessage)

	select {}
}

func handleMessage(msg pkg.ReadMessage) {

	fmt.Printf("topic: %s, Message: %s, Partition: %d, Key: %d\n", msg.Topic, msg.Value, msg.Partition, msg.Key)

	println("Headers:")
	for _, header := range msg.Headers {
		fmt.Printf("Key: %s, Value: %s\n", header.Key, header.Value)
	}

}
