package kafka

import (
	"context"
	"fmt"

	apmkafkago "github.com/rafaelsouzaribeiro/apm-kafkago/pkg"
	"github.com/rafaelsouzaribeiro/consumer/pkg"
	"github.com/segmentio/kafka-go"
)

func (c *Reader) Receive(r *pkg.ReadMessage, canal chan<- pkg.ReadMessage) {
	k := kafka.ReaderConfig{
		Brokers:   c.Brokers,
		Topic:     r.Topic,
		Partition: r.Partition,
	}

	if r.GroupID != "" {
		k.GroupID = r.GroupID
	}

	kReader := kafka.NewReader(k)

	reader := apmkafkago.WrapReader(kReader)

	const batchSize = 10
	var messagesToCommit []kafka.Message

	for {
		msg, err := reader.ReadMessage(context.Background())

		if err != nil {
			fmt.Println("Error reading message:", err)
			continue
		}

		canal <- pkg.ReadMessage{
			Topic:     msg.Topic,
			Value:     string(msg.Value),
			Partition: msg.Partition,
			Key:       msg.Key,
			Time:      msg.Time,
			GroupID:   r.GroupID,
			Headers:   *getHeader(msg),
		}

		messagesToCommit = append(messagesToCommit, msg)

		if len(messagesToCommit) >= batchSize {
			err = reader.R.CommitMessages(context.Background(), messagesToCommit...)

			if err != nil {
				fmt.Println("Error committing messages:", err)
			}

			messagesToCommit = nil
		}
	}
}

func getHeader(msg kafka.Message) *[]pkg.Header {
	var headers []pkg.Header
	for _, header := range msg.Headers {
		headers = append(headers, pkg.Header{
			Key:   header.Key,
			Value: string(header.Value),
		})
	}

	return &headers
}
