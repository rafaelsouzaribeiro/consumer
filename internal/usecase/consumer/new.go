package consumer

import (
	"github.com/rafaelsouzaribeiro/consumer/internal/entity"
)

type ConsumerUseCase struct {
	consumer entity.Iconsumer
}

func NewConsumerUseCase(e entity.Iconsumer) *ConsumerUseCase {
	return &ConsumerUseCase{
		consumer: e,
	}
}
