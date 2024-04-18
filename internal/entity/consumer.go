package entity

import "github.com/rafaelsouzaribeiro/apache-kafka/consumer"

type Consumer struct {
	Brokers           []string
	GroupId           string
	Topics            []string
	HandleMessageFunc consumer.MessageCallback
}

func NewConsumer(brokers []string, groupId string, topics []string, handleMessageFunc consumer.MessageCallback) *Consumer {

	return &Consumer{
		Brokers:           brokers,
		GroupId:           groupId,
		Topics:            topics,
		HandleMessageFunc: handleMessageFunc,
	}
}
