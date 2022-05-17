package consumer

import (
	"log"

	"github.com/Shopify/sarama"
)

func newConsumerConn(brokerList []string) sarama.ConsumerGroup {
	config := sarama.NewConfig()
	client, err := sarama.NewConsumerGroup(brokerList, "ChatService", config)
	if err != nil {
		log.Fatalln("Failed to start Sarama producer:", err)
	}
	log.Printf("connected to brokers %v", brokerList)
	return client
}
