package kafka

import (
	"github.com/rafaelsouzaribeiro/consumer/internal/infra/receive/kafka"
	"github.com/rafaelsouzaribeiro/consumer/internal/usecase/consumer"
)

func NewConsumerUseCase() consumer.ConsumerUseCase {

	kafkaProducer := kafka.NewConsumer()
	myService := consumer.NewConsumerUseCase(kafkaProducer)

	return *myService

}
