package kafka

import (
	"context"
	"fmt"

	apmkafkago "github.com/rafaelsouzaribeiro/apm-kafkago/pkg"
	"github.com/rafaelsouzaribeiro/consumer/pkg"
	"github.com/segmentio/kafka-go"
)

func (c *Reader) Receive(r *pkg.ReadMessage, handleMessage func(ms *pkg.ReadMessage)) {
	// Consumer
	kReader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: c.Brokers,
		Topic:   r.Topic,
		GroupID: r.GroupID,
	})

	reader := apmkafkago.WrapReader(kReader)

	const batchSize = 10
	var messagesToCommit []kafka.Message

	for {
		msg, err := reader.ReadMessage(context.Background())

		if err != nil {
			fmt.Println("Error reading message:", err)
			continue
		}

		co := c.UpdateKafkaMessage(r, &msg)

		handleMessage(&co)

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
