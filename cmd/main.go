package main

import (
	"github.com/rafaelsouzaribeiro/consumer/pkg"
	"github.com/rafaelsouzaribeiro/consumer/pkg/kafka"
)

func main() {
	exec := kafka.NewReader([]string{"springboot:9092"})
	msg := pkg.ReadMessage{
		Topic:   "contact-adm-insert",
		GroupID: "contact",
	}
	go exec.Receive(&msg, handleMessage)

	select {}
}

func handleMessage(messages, topics string) {

	println(topics, messages)

}
