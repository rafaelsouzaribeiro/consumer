package kafka

import (
	"github.com/IBM/sarama"
	"github.com/rafaelsouzaribeiro/apache-kafka/consumer"
	"github.com/rafaelsouzaribeiro/consumer/internal/entity"
)

type Consumer struct {
}

func NewConsumer() *Consumer {
	return &Consumer{}
}

func (c *Consumer) Receive(receive *entity.Consumer) {
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true
	//config.Consumer.Offsets.AutoCommit.Enable = true
	// config.Net.SASL.Enable = true
	// config.Net.SASL.User = "root"
	// config.Net.SASL.Password = "root"

	con := consumer.NewConsumer(receive.Brokers, receive.GroupId, receive.Topics, config, /*func(messages string, topic string) {
			// Processe as mensagens recebidas aqui
			fmt.Println("Message received:", messages, "topic:", topic)
			// Envia o tópico para o canal

		}*/receive.HandleMessageFunc)

	client, err := con.GetConsumer()

	if err != nil {
		panic(err)
	}

	defer func() {
		if err := client.Close(); err != nil {
			panic(err)
		}

	}()

	cancel, err := con.VerifyConsumer(client)
	defer cancel()

	if err != nil {
		panic(err)
	}

	con.VerifyError(client)

}
