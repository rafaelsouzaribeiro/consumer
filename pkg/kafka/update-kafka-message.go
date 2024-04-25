package kafka

import (
	"github.com/rafaelsouzaribeiro/consumer/pkg"
	"github.com/segmentio/kafka-go"
)

func (c *Reader) UpdateKafkaMessage(rm *pkg.ReadMessage, msg kafka.Message) pkg.ReadMessage {
	rm.Value = string(msg.Value)
	rm.Topic = msg.Topic
	rm.Partition = msg.Partition
	rm.Key = msg.Key

	var headers []pkg.Header
	for _, header := range msg.Headers {
		headers = append(headers, pkg.Header{
			Key:   header.Key,
			Value: string(header.Value),
		})
	}
	rm.Headers = headers

	return *rm
}
