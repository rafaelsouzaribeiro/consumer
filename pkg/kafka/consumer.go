package kafka

import (
	"context"
	"fmt"

	apmkafkago "github.com/rafaelsouzaribeiro/apm-kafkago/pkg"
	"github.com/rafaelsouzaribeiro/consumer/pkg"
	"github.com/segmentio/kafka-go"
)

func (c *Reader) Receive(r *pkg.ReadMessage, handleMessage func(messages, topics string)) {
	// Consumer
	kReader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: c.Brokers,
		Topic:   r.Topic,
		GroupID: r.GroupID,
	})

	reader := apmkafkago.WrapReader(kReader)

	for {
		msg, err := reader.ReadMessage(context.Background())

		if err != nil {
			fmt.Println("Error reading message:", err)
			continue
		}

		handleMessage(string(msg.Value), msg.Topic)

		reader.R.CommitMessages(context.Background())
	}
}
