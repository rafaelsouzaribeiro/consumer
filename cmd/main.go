package main

import (
	"fmt"

	"github.com/rafaelsouzaribeiro/consumer/pkg"
	exec "github.com/rafaelsouzaribeiro/consumer/pkg/kafka"
)

func main() {
	consumer := exec.NewReader([]string{"springboot:9092"})
	msg := pkg.ReadMessage{
		Topic:   []string{"contact-adm-insert"},
		GroupID: "contacts",
	}
	go consumer.Receive(&msg, handleMessage)

	select {}
}

func handleMessage(msg *pkg.ReadMessage) {

	fmt.Printf("topic: %s, Message: %s, Partition: %d, Key: %d, time: %s\n", msg.Topic[0], msg.Value, msg.Partition, msg.Key, msg.Time.Format("2006-01-02 15:04:05"))

	println("Headers:")
	for _, header := range msg.Headers {
		fmt.Printf("Key: %s, Value: %s\n", header.Key, header.Value)
	}

}
