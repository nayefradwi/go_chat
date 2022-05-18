package consumer

import (
	"log"

	"github.com/Shopify/sarama"
)

type ConsumerClientKey struct{}

func NewConsumerClient(brokerList []string) sarama.ConsumerGroup {
	config := sarama.NewConfig()
	config.Consumer.Offsets.Initial = sarama.OffsetNewest
	client, err := sarama.NewConsumerGroup(brokerList, "ChatService", config)
	if err != nil {
		log.Fatalln("Failed to start Sarama producer:", err)
	}
	log.Printf("connected to brokers %v", brokerList)
	return client
}
