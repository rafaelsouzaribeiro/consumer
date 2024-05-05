package main

import (
	"fmt"

	"github.com/rafaelsouzaribeiro/consumer/pkg"
	exec "github.com/rafaelsouzaribeiro/consumer/pkg/kafka"
)

func main() {
	consumer := exec.NewBrokers([]string{"springboot:9092"})
	msg := pkg.ReadMessage{
		Topic: "contact-adm-insert",
		//remove the group to listen to only one partition of the topic
		GroupID: "contacts",
		//to set GroupID set the partition to 0
		Partition: 0,
	}

	canal := make(chan pkg.ReadMessage)
	go consumer.Receive(&msg, canal)

	for obj := range canal {
		fmt.Printf("topic: %s, GroupID: %s, Message: %s, Partition: %d, Key: %d, time: %s\n", obj.Topic, obj.GroupID, obj.Value, obj.Partition, obj.Key, obj.Time.Format("2006-01-02 15:04:05"))

		println("Headers:")
		for _, header := range obj.Headers {
			fmt.Printf("Key: %s, Value: %s\n", header.Key, header.Value)
		}
	}
	select {}
}
