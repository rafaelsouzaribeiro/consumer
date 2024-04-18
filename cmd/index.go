package cmd

import (
	"github.com/rafaelsouzaribeiro/apache-kafka/consumer"
	"github.com/rafaelsouzaribeiro/consumer/internal/infra/di/kafka"
)

func Consumer(brokers []string, groupId string, topics []string, handleMessageFunc consumer.MessageCallback) {
	producer := kafka.NewConsumerUseCase()
	producer.Receive(brokers, groupId, topics, handleMessageFunc)

}
