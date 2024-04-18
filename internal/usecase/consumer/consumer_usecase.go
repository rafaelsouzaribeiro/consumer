package consumer

import (
	"github.com/rafaelsouzaribeiro/apache-kafka/consumer"
	"github.com/rafaelsouzaribeiro/consumer/internal/entity"
)

func (c *ConsumerUseCase) Receive(brokers []string, groupId string, topics []string, handleMessageFunc consumer.MessageCallback) {

	receive := entity.NewConsumer(brokers, groupId, topics, handleMessageFunc)
	c.consumer.Receive(receive)

}
